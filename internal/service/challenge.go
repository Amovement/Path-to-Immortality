package service

import (
	"errors"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/types"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"time"
)

type ChallengeService struct {
	ChallengeMap        map[uint]model.Challenge
	MonsterMap          map[uint]model.Monster
	ChallengeMonsterMap map[uint]model.ChallengeMonster
	ChallengeList       []model.Challenge
}

func NewChallengeService() *ChallengeService {
	return &ChallengeService{
		MonsterMap:          repo.GetMonsterMap(),
		ChallengeMap:        repo.GetChallengeMap(),
		ChallengeMonsterMap: repo.GetChallengeMonsterMap(),
		ChallengeList:       repo.GetChallengeList(),
	}
}

// ListChallenge 获取挑战列表
// 返回值:
//   - types.ListChallengeResp: 包含挑战列表的响应结构体
//   - error: 错误信息，始终返回nil
func (s *ChallengeService) ListChallenge() (types.ListChallengeResp, error) {
	var res []types.ListChallengeItem

	user := getLocalUser()
	userLevel := user.Level
	if userLevel == 0 {
		userLevel = 1
	}

	// 遍历所有挑战ID，加载每个挑战的详细信息
	for _, challengeId := range s.ChallengeList {
		if challengeId.LevelLimit-userLevel >= 30 {
			continue
		}
		c, err := s.LoadChallenge(challengeId.ID)
		if err != nil {
			fmt.Printf("[ERROR] load challenge error: %+v\n", err)
			continue
		}
		res = append(res, c)
	}

	return types.ListChallengeResp{
		List: res,
	}, nil
}

// LoadChallenge 根据挑战ID加载挑战信息
// 参数:
//
//	challengeId: 挑战的唯一标识符
//
// 返回值:
//
//	types.ListChallengeItem: 包含怪物列表、奖励和标题的挑战项信息
//	error: 如果找不到对应挑战则返回错误，否则返回nil
func (s *ChallengeService) LoadChallenge(challengeId uint) (types.ListChallengeItem, error) {
	var monsters []types.Monster

	// 从挑战映射中查找指定ID的挑战
	challenge, ok := s.ChallengeMap[challengeId]
	if !ok {
		return types.ListChallengeItem{}, errors.New("challenge not found")
	}

	// 遍历所有挑战怪物映射，找出属于当前挑战的所有怪物
	for _, challengeMonster := range s.ChallengeMonsterMap {
		if challengeMonster.ChallengeID == challengeId {
			// 构造怪物对象，包含基础属性和通过等级计算的修为信息
			monster := types.Monster{
				Name:        challengeMonster.Monster.Name,
				Hp:          challengeMonster.Monster.Hp,
				HpLimit:     challengeMonster.Monster.HpLimit,
				Attack:      challengeMonster.Monster.Attack,
				Defense:     challengeMonster.Monster.Defense,
				Speed:       challengeMonster.Monster.Speed,
				Cultivation: utils.GetCultivationByLevel(int(challengeMonster.Monster.Level)),
			}
			// 根据怪物数量重复添加到怪物列表中
			for cnt := 0; cnt < int(challengeMonster.Count); cnt++ {
				monsters = append(monsters, monster)
			}
		}
	}

	if challengeId == 0 { // 特判每日挑战 - 心魔
		user := getLocalUser()
		monsters = append(monsters, types.Monster{
			Attack:      user.Attack,
			Defense:     user.Defense,
			Hp:          user.HpLimit,
			HpLimit:     user.HpLimit,
			Name:        user.Username + "的心魔",
			Speed:       user.Speed,
			Cultivation: utils.GetCultivationByLevel(int(user.Level)),
		})
	}

	// 构造奖励字符串
	rewards := fmt.Sprintf("金币 %d 枚", challenge.Gold)
	description := "挑战内容: \n" // 怪物群的描述
	for _, monster := range monsters {
		description = description + fmt.Sprintf(" [%s-%s]", monster.Name, monster.Cultivation)
		description = description + fmt.Sprintf("体魄: %d/%d ", monster.Hp, monster.HpLimit)
		description = description + fmt.Sprintf("攻击: %d 防御: %d ", monster.Attack, monster.Defense)
		description = description + fmt.Sprintf("速度: %d \n", monster.Speed)
	}

	// 组装最终的挑战项结果
	res := types.ListChallengeItem{
		MonsterList: monsters,
		Reward:      rewards,
		Title:       challenge.Title,
		ID:          challenge.ID,
		Description: description,
		LevelLimit:  challenge.LevelLimit,
	}

	return res, nil
}

// JoinChallenge 用于让用户加入指定的挑战。
// 参数:
//   - ChallengeId: 要参与的挑战ID
//
// 返回值:
//   - string: 操作结果信息（如战斗结果、错误提示等）
//   - string: 战斗日志内容
func (s *ChallengeService) JoinChallenge(ChallengeId int) (string, string) {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁", ""
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	// 检查是否存在挑战
	challengeCache, existed := s.ChallengeMap[uint(ChallengeId)]
	if !existed {
		return "挑战不存在", ""
	}

	challenge, err := s.LoadChallenge(challengeCache.ID)
	if err != nil {
		return "内部错误...", ""
	}
	user := getLocalUser()

	if user.Exp >= user.Level*10*3 {
		return "你身上的经验积累太多了... 先去消耗一番积蓄再来吧...", "你身上的经验积累太多了... 先去消耗一番积蓄再来吧..."
	}

	if lastPassedTime, existed := user.PassedChallengeTime[challenge.ID]; existed {
		if time.Now().Format("2006-01-02") == lastPassedTime {
			return "今日已经挑战过该项,请明日再来吧...", "今日已经挑战过该项,请明日再来吧..."
		}
	}
	if challenge.LevelLimit-user.Level >= 30 {
		return "修为不足", "修为不足，先去修炼一番吧..."
	}

	msg, fightLog := FightCore(user, challenge.MonsterList)
	if msg == fightWin || msg == fightDraw {
		equipAttr := getUserEquipAttributes()
		if msg == fightWin {
			goldGet := challengeCache.Gold
			if model.CheckHasSpecial(equipAttr.Special, model.SpecialsGreedy) {
				goldGet = (goldGet * 8) / 10
			}
			user.Gold = user.Gold + goldGet
			user = s.userPassChallenge(user, challenge.ID)
			msg = msg + challenge.Title + "战斗胜利. 获得金币 " + fmt.Sprint(goldGet) + " 枚."
		}
		if model.CheckHasSpecial(equipAttr.Special, model.SpecialsBleed) {
			user.Hp -= (user.HpLimit * 5) / 100
			if user.Hp < 0 {
				user.Hp = 1
			}
			msg = msg + " `流血诅咒`扣除了" + fmt.Sprint((user.HpLimit*5)/100) + "点体魄. "
		}
	} else if msg == fightLose {
		// 掉 10% 经验惩罚 + 损失身上 10% 的金币
		// 如果身上经验不足 10% 则直接掉级，所有属性减 1
		if user.Exp <= user.Level { // 升级经验是等级的 10 倍
			user.Level = utils.Max(user.Level-1, 0)
			user.Attack = utils.Max(user.Attack-1, 0)
			user.Defense = utils.Max(user.Defense-1, 0)
			user.Speed = utils.Max(user.Speed-1, 0)
			user.HpLimit = utils.Max(user.HpLimit-1, model.DefaultHp)
			msg = msg + " 经验不足, 您的境界跌落了(各项属性下降)！"
		} else {
			msg = msg + " 修为损失 " + fmt.Sprint(user.Exp/10) + ". "
			user.Exp = user.Exp - user.Exp/10
		}
		msg = msg + " 金币减少 " + fmt.Sprint(user.Gold/10) + ". "
		user.Gold = user.Gold - user.Gold/10

		user.Potential = utils.Max(user.Potential-3, 0) // 潜能也扣
		user.Hp = 1
	}
	updateUserInfo(user)

	return msg, fightLog + "\n" + msg
}

// userPassChallenge 为用户添加已通过的挑战ID
// 参数:
//   - user: 用户对象，用于记录已通过的挑战ID
//   - challengeId: 挑战ID，表示用户通过的挑战
func (s *ChallengeService) userPassChallenge(user *model.User, challengeId uint) *model.User {
	if user.PassedChallengeTime == nil {
		user.PassedChallengeTime = make(map[uint]string)
	}
	user.PassedChallengeTime[challengeId] = time.Now().Format("2006-01-02")
	// 检查用户是否已经通过该挑战，避免重复记录
	for i := 0; i < len(user.PassedChallengeId); i++ {
		if user.PassedChallengeId[i] == challengeId {
			return user
		}
	}
	// 将新的挑战ID添加到用户已通过的挑战列表中
	user.PassedChallengeId = append(user.PassedChallengeId, challengeId)
	return user
}
