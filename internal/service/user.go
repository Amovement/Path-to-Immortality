package service

import (
	"encoding/json"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/types"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"strings"
	"time"
)

type UserService struct {
	challengeList []model.Challenge
}

func NewUserService() *UserService {
	return &UserService{
		challengeList: repo.GetChallengeList(),
	}
}

// getLocalUser 获取本地用户
func getLocalUser() *model.User {
	user := model.NewUser()
	userInfo, existed := utils.GetStorage(utils.UserInfoStorageKey)
	if existed {
		if IsProd() {
			userInfo, _ = utils.Decrypt(userInfo)
		}
		err := json.Unmarshal([]byte(userInfo), &user)
		if err != nil {
			fmt.Printf("[ERROR] %+v\n", err)
			user = model.NewUser()
		}
	}
	return user
}

// updateUserInfo 更新本地用户信息
func updateUserInfo(user *model.User) {
	userBytes, _ := json.Marshal(user)
	userString := string(userBytes)
	if IsProd() {
		userStringEncrypted, err := utils.Encrypt(string(userBytes))
		if err != nil {
			fmt.Printf("[ERROR] %+v\n", err)
			return
		}
		userString = userStringEncrypted
	}
	utils.SetStorage(utils.UserInfoStorageKey, userString)
}

// GetUserInfo 获取本地用户信息
func (s *UserService) GetUserInfo() types.GetUserInfoResp {
	user := getLocalUser()
	ret := types.GetUserInfoResp{
		Username:            user.Username,
		Attack:              user.Attack,
		Defense:             user.Defense,
		Speed:               user.Speed,
		Exp:                 user.Exp,
		Gold:                user.Gold,
		Hp:                  user.Hp,
		HpLimit:             user.HpLimit,
		Potential:           user.Potential,
		Cultivation:         utils.GetCultivationByLevel(int(user.Level)),
		NextCultivationTime: time.Unix(user.NextCultivationTime, 0).Format("2006-01-02 15:04:05"),
		Level:               user.Level,
	}
	if user.RestartCount > 0 {
		ret.Cultivation += fmt.Sprintf(" [转生之人%s]", utils.IntToRoman(user.RestartCount))
	}
	return ret
}

// SetUsername 设置用户用户名
func (s *UserService) SetUsername(username string) {
	// 获取本地用户信息
	user := getLocalUser()
	// 更新用户名
	user.Username = username
	// 保存更新后的用户信息
	updateUserInfo(user)
}

// Allocate 为用户分配属性点
// 参数 stat: 要分配的属性类型，可选值为 "attack"(攻击)、"defense"(防御)、"hpLimit"(生命上限)、"speed"(速度)
// 返回值: 操作结果字符串，成功返回"分配成功"，失败返回相应的错误提示
func (s *UserService) Allocate(stat string) string {
	key := fmt.Sprint("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	user := getLocalUser()

	if user.Potential <= 0 {
		return "潜能点不足"
	}

	// 根据不同的属性类型进行分配
	switch stat {
	case "attack":
		user.Attack = user.Attack + 1
	case "defense":
		user.Defense = user.Defense + 1
	case "hpLimit":
		user.HpLimit = user.HpLimit + 5
	case "speed":
		user.Speed = user.Speed + 1
	default:
		return "属性不存在"
	}
	user.Potential = user.Potential - 1
	updateUserInfo(user)

	return "分配成功"
}

// Heal 处理用户治疗逻辑
// 该函数用于恢复用户血量，治疗和修炼共用时间限制
// 返回值：治疗结果的描述信息
func (s *UserService) Heal() string {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	user := getLocalUser()
	if user.Hp >= user.HpLimit {
		return "当前已经很健康了"
	}

	endTime := user.NextCultivationTime // 治疗和修炼占用同一个时间
	// 需要检查一下时间能不能治疗
	if time.Now().Unix() <= endTime {
		return "累了, 暂时无法进行这个操作. 下一次可操作时间: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
	}

	// 设置下次可操作时间为5分钟后
	endTime = time.Now().Add(time.Minute * 5).Unix()
	user.NextCultivationTime = endTime

	// 恢复血量上限的30%，但不超过血量上限
	user.Hp = user.Hp + (user.HpLimit/10)*3
	if user.Hp > user.HpLimit {
		user.Hp = user.HpLimit
	}
	updateUserInfo(user)

	return "恢复成功! 下一次可修炼/恢复/做工时间点: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
}

func (s *UserService) Cultivation() string {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	msg := ""
	user := getLocalUser()

	endTime := user.NextCultivationTime
	// 需要检查一下时间能不能修炼
	if time.Now().Unix() <= endTime {
		return "累了, 暂时无法进行这个操作. 下一次可操作时间: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
	}

	endTime = time.Now().Add(utils.GetRandomMinutes(1, int(user.Level))).Unix()
	user.NextCultivationTime = endTime
	user.Exp = user.Exp + utils.GetRandomInt64(1, 5)
	if user.Exp >= user.Level*10 { // 升级嘞
		// 检查用户是否达到破境要求
		if userIsBrokenLevel(user) {
			// 破境需要完成对应等级的挑战
			ok := s.checkUserPassedChallenge(user)
			if !ok {
				return "体内的力量积蓄了太多了，请完成对应的圆满挑战后再来突破吧...下一次可修炼/恢复/做工时间点: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
			}
		}
		user.Exp = user.Exp - user.Level*10
		user.Level = user.Level + 1
		user.Potential = user.Potential + 3
		user.HpLimit = user.HpLimit + 5
		user.Hp = user.HpLimit
		msg += "你感觉浑身充满了力量，获得了 3 点潜能，境界提升了..."
		if user.RestartCount > 0 {
			user.Potential = user.Potential + user.RestartCount
			msg += "体内的另外一个灵魂正在回应你, 你好像想起来了很多东西, 额外获得了 " + fmt.Sprint(user.RestartCount) + " 点潜能"
		}
	}
	// 更新用户
	updateUserInfo(user)

	return msg + " 修炼成功! 下一次可修炼/恢复/做工时间点: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
}

// GetGold 打工赚钱
func (s *UserService) GetGold() string {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	msg := ""
	user := getLocalUser()

	endTime := user.NextCultivationTime
	// 需要检查一下时间能不能赚钱
	if time.Now().Unix() <= endTime {
		return "累了, 暂时无法进行这个操作. 下一次可操作时间: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
	}

	endTime = time.Now().Add(utils.GetRandomMinutes(1, int(user.Level))).Unix()
	user.NextCultivationTime = endTime
	getGold := 10 + utils.GetRandomInt64(1, 10)
	user.Gold += getGold
	msg = "恭喜你, 获得金币: " + fmt.Sprint(getGold) + " 枚"
	// 更新用户
	updateUserInfo(user)
	return msg + " 下一次可修炼/恢复/做工时间点: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
}

// userIsBrokenLevel 判断用户等级是否为破境等级
//
//	破境等级定义为：等级值能被30整除且大于0的等级
//
// 参数：
//
//	user: 用户对象指针，包含用户的等级信息
//
// 返回值：
//
//	bool: 如果用户等级是破损等级返回true，否则返回false
func userIsBrokenLevel(user *model.User) bool {
	// 判断等级是否能被30整除且大于0
	if user.Level%30 == 0 && user.Level > 0 {
		return true
	}
	return false
}

// checkUserPassedChallenge 检查用户是否通过了对应等级的圆满挑战
// 参数:
//   - user: 用户对象，包含用户等级和已通过的挑战ID列表
//
// 返回值:
//   - bool: 如果用户通过了对应等级的圆满挑战返回true，否则返回false
func (s *UserService) checkUserPassedChallenge(user *model.User) bool {
	// 构建用户已通过挑战的映射表，便于快速查找
	userPassedMap := make(map[uint]bool)
	for _, challengeId := range user.PassedChallengeId {
		userPassedMap[challengeId] = true
	}
	userLevel := user.Level

	// 遍历挑战列表，查找与用户等级匹配且标题包含"圆满挑战"的挑战
	for _, challenge := range s.challengeList {
		if challenge.LevelLimit == userLevel && strings.Contains(challenge.Title, "圆满挑战") {
			// 找到对应的圆满挑战
			if userPassedMap[challenge.ID] {
				// 玩家已通过
				return true
			}
			return false
		}
	}
	return false
}

// Restart 重置 轮回转生
func (s *UserService) Restart() string {
	var msg string
	user := getLocalUser()
	newUser := model.NewUser()
	newUser.Username = user.Username
	newUser.RestartCount = user.RestartCount

	if user.Level >= 150 {
		newUser.RestartCount += 1
		msg = "一阵恍惚后...你重新睁开了眼睛, 年轻的肉体里充满了使不完的劲. 你意识到自己已经转生 " + fmt.Sprint(newUser.RestartCount) + " 次. "
	} else {
		msg = "一阵恍惚后...你重新睁开了眼睛, 你好像忘记了很多事情... "
	}

	updateUserInfo(newUser)
	return msg
}
