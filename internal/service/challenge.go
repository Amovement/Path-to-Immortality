package service

import (
	"errors"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/types"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"math/rand"
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

	// 遍历所有挑战ID，加载每个挑战的详细信息
	for _, challengeId := range s.ChallengeList {
		if user.Level > challengeId.LevelLimit || challengeId.LevelLimit-user.Level >= 30 {
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

	if user.Level > challenge.LevelLimit {
		return "不要以大欺小啊", "不要以大欺小啊"
	}
	if challenge.LevelLimit-user.Level >= 30 {
		return "修为不足", "修为不足"
	}

	msg, fightLog := s.fightCore(user, challenge.MonsterList)
	if msg == fightWin {
		user.Gold = user.Gold + challengeCache.Gold
		msg = msg + " 获得金币 " + fmt.Sprint(challengeCache.Gold) + " 枚"
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

	return msg, fightLog
}

const (
	fightWin  = "战斗胜利!"
	fightLose = "战斗失败!"
	// 暴击倍率（修真世界 "神通触发" 概率）
	critRate  = 0.15 // 15% 暴击概率
	critMulti = 1.8  // 暴击伤害倍率
	// 闪避概率（修真世界 "身法闪避" 概率）
	dodgeRate = 0.08 // 8% 闪避概率
)

// fightCore 战斗核心逻辑：玩家 VS 多怪物回合制战斗，按速度排序决定出手顺序
//
//	返回战斗结果和战斗日志
func (s *ChallengeService) fightCore(user *model.User, monsters []types.Monster) (string, string) {
	// 1. 初始化战斗状态（深拷贝避免修改原数据）
	player := copyUserState(user)
	monsterStates := copyMonsterStates(monsters)
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子（确保每次战斗结果不同）
	logMsg := "玩家 " + player.Username + " 和 "
	for _, monster := range monsterStates {
		logMsg += " " + monster.Name
	}
	logMsg += " 战斗开始了！\n"

	// 2. 战斗主循环（直到玩家死亡或所有怪物死亡）
	for {
		// 检查战斗结束条件
		if isPlayerDead(player) {
			return fightLose, logMsg
		}
		if isAllMonstersDead(monsterStates) {
			// 更新玩家剩余血量
			user.Hp = player.Hp
			return fightWin, logMsg
		}

		// 3. 计算当前回合所有战斗单位的出手顺序（按速度排序，速度相同则玩家优先）
		actionOrder := getActionOrder(player, monsterStates)

		// 4. 执行当前回合所有单位的攻击动作
		for _, unit := range actionOrder {
			if unit.isPlayer {
				// 玩家出手：选择一个存活怪物攻击
				targetMonster := getRandomAliveMonster(monsterStates)
				if targetMonster != nil {
					logMsg += playerAttack(player, targetMonster)
				}
			} else {
				// 怪物出手：攻击玩家（若怪物已死亡则跳过）
				monster := unit.monster
				if monster.Hp <= 0 {
					continue
				}
				logMsg += monsterAttack(monster, player)
			}

			// 攻击后立即检查战斗结束（避免后续单位无效出手）
			if isPlayerDead(player) {
				return fightLose, logMsg
			}
			if isAllMonstersDead(monsterStates) {
				// 更新玩家剩余血量
				user.Hp = player.Hp
				return fightWin, logMsg
			}
		}
	}
}

//--------------- 战斗辅助函数 ---------------

// copyUserState 深拷贝玩家状态（避免战斗中修改原用户数据）
func copyUserState(user *model.User) *model.User {
	return &model.User{
		Attack:  user.Attack,
		Defense: user.Defense,
		Hp:      user.Hp,
		HpLimit: user.HpLimit,
		Speed:   user.Speed,
		Level:   user.Level,
	}
}

// copyMonsterStates 深拷贝怪物状态（处理多怪物战斗状态）
func copyMonsterStates(monsters []types.Monster) []*types.Monster {
	var states []*types.Monster
	for _, m := range monsters {
		states = append(states, &types.Monster{
			Name:        m.Name,
			Hp:          m.Hp,
			HpLimit:     m.HpLimit,
			Attack:      m.Attack,
			Defense:     m.Defense,
			Speed:       m.Speed,
			Cultivation: m.Cultivation,
		})
	}
	return states
}

// isPlayerDead 检查玩家是否死亡（Hp≤0 判定死亡）
func isPlayerDead(player *model.User) bool {
	return player.Hp <= 0
}

// isAllMonstersDead 检查所有怪物是否死亡
func isAllMonstersDead(monsters []*types.Monster) bool {
	for _, m := range monsters {
		if m.Hp > 0 {
			return false
		}
	}
	return true
}

// actionUnit 战斗单位动作排序结构体
type actionUnit struct {
	isPlayer bool           // 是否为玩家
	speed    int64          // 速度（决定出手顺序）
	monster  *types.Monster // 怪物指针（非玩家时有效）
}

// getActionOrder 计算当前回合出手顺序：速度降序，速度相同则玩家优先
func getActionOrder(player *model.User, monsters []*types.Monster) []actionUnit {
	var units []actionUnit
	// 添加玩家到动作列表
	units = append(units, actionUnit{
		isPlayer: true,
		speed:    player.Speed,
	})
	// 添加所有存活怪物到动作列表
	for _, m := range monsters {
		if m.Hp > 0 {
			units = append(units, actionUnit{
				isPlayer: false,
				speed:    m.Speed,
				monster:  m,
			})
		}
	}

	// 按速度排序（降序），速度相同则玩家在前
	for i := 0; i < len(units); i++ {
		for j := i + 1; j < len(units); j++ {
			if units[j].speed > units[i].speed {
				units[i], units[j] = units[j], units[i]
			} else if units[j].speed == units[i].speed && units[j].isPlayer {
				units[i], units[j] = units[j], units[i]
			}
		}
	}
	return units
}

// getRandomAliveMonster 随机选择一个存活的怪物作为攻击目标
func getRandomAliveMonster(monsters []*types.Monster) *types.Monster {
	var aliveMonsters []*types.Monster
	for _, m := range monsters {
		if m.Hp > 0 {
			aliveMonsters = append(aliveMonsters, m)
		}
	}
	if len(aliveMonsters) == 0 {
		return nil
	}
	// 随机选择目标（模拟玩家 "自主选择目标" 的随机性）
	return aliveMonsters[rand.Intn(len(aliveMonsters))]
}

// playerAttack 处理玩家对怪物的攻击逻辑，计算伤害并更新怪物血量
// 参数:
//
//	player: 玩家对象，包含攻击属性
//	monster: 怪物对象，包含防御、血量等属性
//
// 返回值:
//
//	string: 攻击过程的日志信息
func playerAttack(player *model.User, monster *types.Monster) string {
	var logMsg string
	// 1. 计算基础伤害（玩家攻击 - 怪物防御，最低 1 点伤害）
	baseDmg := player.Attack - monster.Defense
	if baseDmg < 1 {
		baseDmg = 1
	}

	// 2. 判定是否触发暴击（修真 "神通暴击"）
	var finalDmg int64
	if rand.Float64() <= critRate {
		finalDmg = int64(float64(baseDmg) * critMulti)
		// 可在此处添加暴击日志：fmt.Printf ("玩家触发神通暴击！对【% s】造成 %.0f 点伤害 \n", monster.Name, critMulti)
		logMsg += fmt.Sprintf("玩家触发神通暴击！对【%s】造成 %.0f 点伤害 \n", monster.Name, critMulti)
	} else {
		finalDmg = baseDmg
		// 可在此处添加普通攻击日志：fmt.Printf ("玩家对【% s】造成 % d 点伤害 \n", monster.Name, baseDmg)
		logMsg += fmt.Sprintf("玩家对【%s】造成 % d 点伤害 \n", monster.Name, baseDmg)
	}

	// 3. 计算怪物剩余血量（最低 0 点）
	monster.Hp -= finalDmg
	if monster.Hp < 0 {
		monster.Hp = 0
	}
	// 可在此处添加伤害结果日志：fmt.Printf ("【% s】剩余血量：% d/% d\n", monster.Name, monster.Hp, monster.HpLimit)
	logMsg += fmt.Sprintf("【%s】剩余血量：% d/% d\n", monster.Name, monster.Hp, monster.HpLimit)
	return logMsg
}

// monsterAttack 处理怪物对玩家的攻击逻辑，包括闪避、暴击、伤害计算和血量更新。
// 参数:
//   - monster: 攻击玩家的怪物对象，包含攻击值等属性
//   - player: 被攻击的玩家对象，包含防御值、当前血量等属性
//
// 返回值:
//   - string: 战斗过程的日志信息，用于展示给用户
func monsterAttack(monster *types.Monster, player *model.User) string {
	var logMsg string

	// 判定是否触发闪避（修真 "身法闪避"）
	if rand.Float64() <= dodgeRate {
		logMsg += fmt.Sprintf("玩家触发身法闪避！躲开了【%s】的攻击 \n", monster.Name)
		return logMsg
	}

	// 计算基础伤害（怪物攻击 - 玩家防御，最低 1 点伤害）
	baseDmg := monster.Attack - player.Defense
	if baseDmg < 1 {
		baseDmg = 1
	}

	// 判定是否触发怪物暴击（高阶怪物 "妖术暴击"，概率比玩家低 5%）
	monsterCritRate := critRate - 0.05
	if monsterCritRate < 0 {
		monsterCritRate = 0
	}
	var finalDmg int64
	if rand.Float64() <= monsterCritRate {
		finalDmg = int64(float64(baseDmg) * critMulti)
		logMsg += fmt.Sprintf("【%s】触发妖术暴击！对玩家造成 %.0f 点伤害 \n", monster.Name, critMulti)
	} else {
		finalDmg = baseDmg
		logMsg += fmt.Sprintf("【%s】对玩家造成 % d 点伤害 \n", monster.Name, baseDmg)
	}

	// 计算玩家剩余血量（最低 0 点）
	player.Hp -= finalDmg
	if player.Hp < 0 {
		player.Hp = 0
	}
	logMsg += fmt.Sprintf("玩家剩余血量：% d/% d\n", player.Hp, player.HpLimit)
	return logMsg
}
