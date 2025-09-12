package service

import (
	"encoding/json"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/types"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// getLocalUser 获取本地用户
func getLocalUser() *model.User {
	user := model.NewUser()
	userInfo, existed := utils.GetStorage(utils.UserInfoStorageKey)
	if existed {
		userInfo, _ = utils.Decrypt(userInfo)
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
	userStringEncrypted, err := utils.Encrypt(string(userBytes))
	if err != nil {
		fmt.Printf("[ERROR] %+v\n", err)
		return
	}
	utils.SetStorage(utils.UserInfoStorageKey, userStringEncrypted)
}

// GetUserInfo 获取本地用户信息
func (s *UserService) GetUserInfo() types.GetUserInfoResp {
	user := getLocalUser()
	return types.GetUserInfoResp{
		Username:            user.Username,
		Attack:              user.Attack,
		Defense:             user.Defense,
		Exp:                 user.Exp,
		Gold:                user.Gold,
		Hp:                  user.Hp,
		HpLimit:             user.HpLimit,
		Potential:           user.Potential,
		Cultivation:         utils.GetCultivationByLevel(int(user.Level)),
		NextCultivationTime: time.Unix(user.NextCultivationTime, 0).Format("2006-01-02 15:04:05"),
	}
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
		return "stat 不存在"
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

	endTime := user.NextCultivationTime // 治疗和修炼占用同一个时间
	// 需要检查一下时间能不能治疗
	if time.Now().Unix() <= endTime {
		return "累了, 暂时无法进行这个操作"
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

	return "恢复成功"
}

func (s *UserService) Cultivation() string {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	user := getLocalUser()

	endTime := user.NextCultivationTime
	// 需要检查一下时间能不能修炼
	if time.Now().Unix() <= endTime {
		return "累了, 暂时无法进行这个操作"
	}

	endTime = time.Now().Add(utils.GetRandomMinutes(1, int(user.Level))).Unix()
	user.NextCultivationTime = endTime
	user.Exp = user.Exp + utils.GetRandomInt64(1, 5)
	if user.Exp >= user.Level*10 { // 升级嘞
		user.Exp = user.Exp - user.Level*10
		user.Level = user.Level + 1
		user.Potential = user.Potential + 3
		user.HpLimit = user.HpLimit + 5
		user.Hp = user.HpLimit
	}
	// 更新用户
	updateUserInfo(user)

	return "下一次可修炼时间点: " + time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
}
