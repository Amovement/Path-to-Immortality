package service

import (
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/types"
	"math/rand"
	"time"
)

const (
	fightWin  = "战斗胜利!"
	fightLose = "战斗失败!"
	fightDraw = "战斗平局!"
	// 暴击倍率（修真世界 "神通触发" 概率）
	critRate  = 0.15 // 15% 暴击概率
	critMulti = 1.8  // 暴击伤害倍率
	// 闪避概率（修真世界 "身法闪避" 概率）
	dodgeRate = 0.08 // 8% 闪避概率
)

// FightCore 战斗核心逻辑：玩家 VS 多怪物回合制战斗，按速度排序决定出手顺序
//
//	返回战斗结果和战斗日志
func FightCore(user *model.User, monsters []types.Monster) (string, string) {
	equipAttr := getUserEquipAttributes() // 获取装备
	// 1. 初始化战斗状态（深拷贝避免修改原数据）
	player := copyUserState(user, &equipAttr)
	monsterStates := copyMonsterStates(monsters)
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子（确保每次战斗结果不同）
	logMsg := "玩家 " + player.Username + " 和 "
	for _, monster := range monsterStates {
		logMsg += " " + monster.Name
	}
	logMsg += " 战斗开始了！\n"

	var nowRoundCount int64
	// 2. 战斗主循环（直到玩家死亡或所有怪物死亡）
	for {
		nowRoundCount++
		if nowRoundCount > 100 {
			logMsg += " 战斗时间过长，结果为平局！\n"
			user.Hp = player.Hp
			return fightDraw, logMsg
		}
		// 检查战斗结束条件
		if isPlayerDead(player) {
			logMsg += "玩家 " + player.Username + " 输掉了战斗！\n"
			return fightLose, logMsg
		}
		if isAllMonstersDead(monsterStates) {
			// 更新玩家剩余体魄
			user.Hp = player.Hp
			logMsg += "玩家 " + player.Username + " 战斗胜利！\n"
			return fightWin, logMsg
		}

		// 3. 计算当前回合所有战斗单位的出手顺序（按速度排序，速度相同则玩家优先）
		actionOrder := getActionOrder(player, equipAttr.Special, monsterStates)

		// 4. 执行当前回合所有单位的攻击动作
		for _, unit := range actionOrder {
			if unit.isPlayer {
				// 玩家出手：选择一个存活怪物攻击
				targetMonster := getRandomAliveMonster(monsterStates)
				if targetMonster != nil {
					logMsg += playerAttack(player, equipAttr.Special, targetMonster)
				}
				if model.CheckHasSpecial(equipAttr.Special, model.SpecialsFast) {
					logMsg += "`迅捷`，使速度提升 " + fmt.Sprint(int64(float64(player.Speed)*0.025)) + " 点.\n"
					player.Speed = player.Speed + int64(float64(player.Speed)*0.025)
				}
				if model.CheckHasSpecial(equipAttr.Special, model.SpecialsSuperFast) {
					logMsg += "`超级迅捷`，使速度提升 " + fmt.Sprint(int64(float64(player.Speed)*0.05)) + " 点.\n"
					player.Speed = player.Speed + int64(float64(player.Speed)*0.05)
				}
			} else {
				// 怪物出手：攻击玩家（若怪物已死亡则跳过）
				monster := unit.monster
				if monster.Hp <= 0 {
					continue
				}
				logMsg += monsterAttack(monster, player, equipAttr.Special)
			}

			// 攻击后立即检查战斗结束（避免后续单位无效出手）
			if isPlayerDead(player) {
				logMsg += "玩家 " + player.Username + " 输掉了战斗！\n"
				return fightLose, logMsg
			}
			if isAllMonstersDead(monsterStates) {
				// 更新玩家剩余体魄
				user.Hp = player.Hp
				logMsg += "玩家 " + player.Username + " 战斗胜利！\n"
				return fightWin, logMsg
			}
		}
	}
}

//--------------- 战斗辅助函数 ---------------

// copyUserState 深拷贝玩家状态（避免战斗中修改原用户数据）
func copyUserState(user *model.User, equipAttr *model.Equip) *model.User {
	return &model.User{
		Username: user.Username,
		Attack:   user.Attack + equipAttr.Attack,
		Defense:  user.Defense + equipAttr.Defense,
		Hp:       user.Hp,
		HpLimit:  user.HpLimit + equipAttr.Hp,
		Speed:    user.Speed + equipAttr.Speed,
		Level:    user.Level,
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
			Special:     m.Special,
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
func getActionOrder(player *model.User, specials []string, monsters []*types.Monster) []actionUnit {
	var units []actionUnit
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
	// 添加玩家到动作列表
	units = append(units, actionUnit{
		isPlayer: true,
		speed:    player.Speed,
	})

	if model.CheckHasSpecial(specials, model.SpecialsSlow) {
		if rand.Float64() <= 0.3 {
			return units
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

// playerAttack 处理玩家对怪物的攻击逻辑，计算伤害并更新怪物体魄
// 参数:
//
//	player: 玩家对象，包含攻击属性
//	specials: 所有装备的特效
//	monster: 怪物对象，包含防御、体魄等属性
//
// 返回值:
//
//	string: 攻击过程的日志信息
func playerAttack(player *model.User, specials []string, monster *types.Monster) string {
	var logMsg string

	if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterEthereal) {
		if rand.Float64() <= 0.3 {
			logMsg += fmt.Sprintf("【%s】`虚体`触发了，闪避了攻击 \n", monster.Name)
			return logMsg
		}
	}

	// 1. 计算基础伤害（玩家攻击 - 怪物防御，最低 1 点伤害）
	baseDmg := player.Attack - monster.Defense
	if baseDmg < 1 {
		baseDmg = 1
	}

	playerCritRate := critRate
	if model.CheckHasSpecial(specials, model.SpecialsCritical) {
		playerCritRate += 0.1
	}
	if model.CheckHasSpecial(specials, model.SpecialsSuperCritical) {
		playerCritRate += 0.2
	}
	playerCritMulti := critMulti
	if model.CheckHasSpecial(specials, model.SpecialsMastery) {
		playerCritMulti += 0.25
	}
	if model.CheckHasSpecial(specials, model.SpecialsSuperMastery) {
		playerCritMulti += 0.5
	}

	// 2. 判定是否触发暴击（修真 "神通暴击"）
	var finalDmg int64
	if rand.Float64() <= playerCritRate {
		finalDmg = int64(float64(baseDmg) * playerCritMulti)
		logMsg += fmt.Sprintf("玩家触发神通暴击！对【%s】造成 %d 点伤害", monster.Name, finalDmg)
	} else {
		finalDmg = baseDmg
		logMsg += fmt.Sprintf("玩家对【%s】造成 %d 点伤害", monster.Name, baseDmg)
	}
	if player.Hp < player.HpLimit && model.CheckHasSpecial(specials, model.SpecialsSuckBlood) {
		player.Hp += finalDmg / 10
		if player.Hp > player.HpLimit {
			player.Hp = player.HpLimit
		}
		logMsg += fmt.Sprintf("触发`体魄偷取`体魄恢复了 %d 点 ", finalDmg/10)
	}
	if model.CheckHasSpecial(specials, model.SpecialsSuperSuckBlood) {
		player.Hp += finalDmg / 5
		if player.Hp > player.HpLimit {
			player.Hp = player.HpLimit
		}
		logMsg += fmt.Sprintf("触发`超级体魄偷取`体魄恢复了 %d 点 ", finalDmg/5)
	}

	if model.CheckHasSpecial(specials, model.SpecialsSharp) {
		finalDmg += (player.Attack * 5) / 100
		logMsg += fmt.Sprintf("额外造成了`尖锐伤害` %d 点 ", (player.Attack*5)/100)
	}
	if model.CheckHasSpecial(specials, model.SpecialsSuperSharp) {
		finalDmg += player.Attack / 10
		logMsg += fmt.Sprintf("额外造成了`超级尖锐伤害` %d 点 ", player.Attack/10)
	}

	// 3. 计算怪物剩余体魄（最低 0 点）
	monster.Hp -= finalDmg
	if monster.Hp <= 0 {
		monster.Hp = 0
		if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterVengeful) {
			player.Hp -= (monster.Attack * 7) / 10
			logMsg += fmt.Sprintf(" 【%s】`复仇`造成 %d 伤害 ", monster.Name, (monster.Attack*7)/10)
		}
		logMsg += fmt.Sprintf(" 【%s】已死亡！\n", monster.Name)
		return logMsg
	}
	if monster.Hp < (monster.HpLimit/10) && model.CheckHasSpecial(monster.Special, model.SpecialsMonsterVolatile) {
		player.Hp -= monster.Hp
		logMsg += fmt.Sprintf(" 【%s】`易爆`造成 %d 伤害,【%s】 已死亡", monster.Name, monster.Hp, monster.Name)
		monster.Hp = 0
		return logMsg
	}

	if monster.Hp > 0 && model.CheckHasSpecial(monster.Special, model.SpecialsMonsterTough) {
		monster.Hp += monster.HpLimit / 10
		if monster.Hp > monster.HpLimit {
			monster.Hp = monster.HpLimit
		}
		logMsg += fmt.Sprintf("【%s】触发`强韧`恢复了 %d 点体魄 ", monster.Name, monster.HpLimit/10)
	}
	if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterThorns) {
		player.Hp -= finalDmg / 10
		logMsg += fmt.Sprintf("【%s】`荆棘`造成 %d 点伤害 ", monster.Name, finalDmg/10)
	}
	logMsg += fmt.Sprintf(" 【%s】剩余体魄：%d/%d\n", monster.Name, monster.Hp, monster.HpLimit)
	if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterShield) {
		logMsg += fmt.Sprintf("【%s】`屏障`增加了 %d 点防御 ", monster.Name, (monster.Defense*5)/100)
		monster.Defense += (monster.Defense * 5) / 100
	}
	if player.Hp <= 0 {
		return logMsg
	}

	if player.Speed > monster.Speed*2 { // 速度超出两倍会触发多次伤害
		for i := 1; i <= int(player.Speed/(monster.Speed+1)); i++ {
			finalDmg = baseDmg
			if rand.Float64() <= playerCritRate {
				finalDmg = int64(float64(baseDmg) * playerCritMulti)
			}
			logMsg += fmt.Sprintf("玩家触发追击对【%s】造成 %d 点伤害 ", monster.Name, finalDmg)
			if player.Hp < player.HpLimit && model.CheckHasSpecial(specials, model.SpecialsSuckBlood) {
				player.Hp += (finalDmg / 10)
				if player.Hp > player.HpLimit {
					player.Hp = player.HpLimit
				}
				logMsg += fmt.Sprintf("触发`体魄偷取`体魄恢复了 %d 点 ", finalDmg/10)
			}
			if model.CheckHasSpecial(specials, model.SpecialsSuperSuckBlood) {
				player.Hp += finalDmg / 5
				if player.Hp > player.HpLimit {
					player.Hp = player.HpLimit
				}
				logMsg += fmt.Sprintf("触发`超级体魄偷取`体魄恢复了 %d 点 ", finalDmg/5)
			}
			if model.CheckHasSpecial(specials, model.SpecialsSharp) {
				finalDmg += (player.Attack * 5) / 100
				logMsg += fmt.Sprintf("额外造成了`尖锐伤害` %d 点 ", (player.Attack*5)/100)
			}
			if model.CheckHasSpecial(specials, model.SpecialsSuperSharp) {
				finalDmg += player.Attack / 10
				logMsg += fmt.Sprintf("额外造成了`超级尖锐伤害` %d 点 ", player.Attack/10)
			}
			monster.Hp -= finalDmg
			if monster.Hp < 0 {
				monster.Hp = 0
			}
			if monster.Hp <= 0 {
				logMsg += fmt.Sprintf("【%s】已死亡！\n", monster.Name)
				if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterVengeful) {
					player.Hp -= (monster.Attack * 7) / 10
					logMsg += fmt.Sprintf(" 【%s】`复仇`造成 %d 伤害 ", monster.Name, (monster.Attack*7)/10)
				}
				break
			}
			if monster.Hp < (monster.HpLimit/10) && model.CheckHasSpecial(monster.Special, model.SpecialsMonsterVolatile) {
				player.Hp -= monster.Hp
				logMsg += fmt.Sprintf(" 【%s】`易爆`造成 %d 伤害,【%s】 已死亡", monster.Name, monster.Hp, monster.Name)
				monster.Hp = 0
				return logMsg
			}
			if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterShield) {
				logMsg += fmt.Sprintf("【%s】`屏障`增加了 %d 点防御 ", monster.Name, (monster.Defense*5)/100)
				monster.Defense += (monster.Defense * 5) / 100
			}
			if monster.Hp > 0 && model.CheckHasSpecial(monster.Special, model.SpecialsMonsterTough) {
				monster.Hp += monster.HpLimit / 10
				if monster.Hp > monster.HpLimit {
					monster.Hp = monster.HpLimit
				}
				logMsg += fmt.Sprintf("【%s】触发`强韧`恢复了 %d 点体魄 ", monster.Name, monster.HpLimit/10)
			}
			if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterThorns) {
				player.Hp -= finalDmg / 10
				logMsg += fmt.Sprintf("【%s】`荆棘`造成 %d 点伤害 ", monster.Name, finalDmg/10)
			}
			logMsg += fmt.Sprintf("【%s】剩余体魄：%d/%d\n", monster.Name, monster.Hp, monster.HpLimit)
		}
	}

	return logMsg
}

// monsterAttack 处理怪物对玩家的攻击逻辑，包括闪避、暴击、伤害计算和体魄更新。
// 参数:
//   - monster: 攻击玩家的怪物对象，包含攻击值等属性
//   - player: 被攻击的玩家对象，包含防御值、当前体魄等属性
//   - specials: 所有装备的特效
//
// 返回值:
//   - string: 战斗过程的日志信息，用于展示给用户
func monsterAttack(monster *types.Monster, player *model.User, specials []string) string {
	var logMsg string

	playerDodgeRate := dodgeRate
	if model.CheckHasSpecial(specials, model.SpecialsSpeedUp) {
		playerDodgeRate += 0.1
	}
	if model.CheckHasSpecial(specials, model.SpecialsSuperSpeedUp) {
		playerDodgeRate -= 0.2
	}
	if model.CheckHasSpecial(specials, model.SpecialsNoob) {
		playerDodgeRate = -1.0
	}

	// 判定是否触发闪避（修真 "身法闪避"）
	if rand.Float64() <= playerDodgeRate {
		logMsg += fmt.Sprintf("玩家触发身法闪避！躲开了【%s】的攻击 \n", monster.Name)
		return logMsg
	}

	// 计算基础伤害（怪物攻击 - 玩家防御，最低 1 点伤害）
	baseDmg := monster.Attack - player.Defense
	if baseDmg < 1 {
		baseDmg = 1
	}

	// 判定是否触发怪物暴击
	monsterCritRate := critRate
	if model.CheckHasSpecial(specials, model.SpecialsWeak) {
		monsterCritRate += 0.2
	}
	if monsterCritRate < 0 {
		monsterCritRate = 0
	}
	var finalDmg int64
	if rand.Float64() <= monsterCritRate {
		finalDmg = int64(float64(baseDmg) * critMulti)
		logMsg += fmt.Sprintf("【%s】触发妖术暴击！对玩家造成 %d 点伤害 ", monster.Name, finalDmg)
	} else {
		finalDmg = baseDmg
		logMsg += fmt.Sprintf("【%s】对玩家造成 %d 点伤害 ", monster.Name, baseDmg)
	}
	if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterRuthless) {
		finalDmg += (finalDmg * 30) / 100
		logMsg += fmt.Sprintf("`残暴`增伤 %d 点伤害 ", (finalDmg*30)/100)
	}
	if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterHuntBlood) {
		logMsg += fmt.Sprintf("【%s】`嗜血`增加了攻击 %d\n ", monster.Name, (monster.Attack*3)/100)
		monster.Attack += (monster.Attack * 3) / 100
	}

	if model.CheckHasSpecial(specials, model.SpecialsSolid) {
		logMsg += fmt.Sprintf("`坚固`抵挡了 %d 点伤害 ", (finalDmg*5)/100)
		finalDmg -= (finalDmg * 5) / 100
	}
	if model.CheckHasSpecial(specials, model.SpecialsSuperSolid) {
		logMsg += fmt.Sprintf("`超级坚固`抵挡了 %d 点伤害 ", (finalDmg*10)/100)
		finalDmg -= (finalDmg * 10) / 100
	}
	if finalDmg <= 0 {
		finalDmg = 1
	}
	if model.CheckHasSpecial(specials, model.SpecialsAggressive) {
		logMsg += fmt.Sprintf("`傲慢诅咒`增加了 %d 点收到的伤害 ", (finalDmg*15)/100)
		finalDmg += (finalDmg * 15) / 100
	}

	// 计算玩家剩余体魄（最低 0 点）
	player.Hp -= finalDmg
	if player.Hp > 0 && model.CheckHasSpecial(specials, model.SpecialsStrong) {
		player.Hp += (player.HpLimit * 3) / 100
		logMsg += fmt.Sprintf("`强壮`恢复了 %d 点体魄 ", (player.HpLimit*3)/100)
	}
	if player.Hp > 0 && model.CheckHasSpecial(specials, model.SpecialsSuperStrong) {
		player.Hp += (player.HpLimit * 6) / 100
		logMsg += fmt.Sprintf("`超级强壮`恢复了 %d 点体魄 ", (player.HpLimit*6)/100)
	}
	if model.CheckHasSpecial(specials, model.SpecialsAchillesHeel) {
		if rand.Float64() <= 0.03 {
			player.Hp = -1
			logMsg += fmt.Sprint("你感觉膝盖中了一箭 `要害诅咒`触发了!!!! 体魄清零了! ")
		}
	}
	logMsg += fmt.Sprintf("玩家剩余体魄：%d/%d\n", player.Hp, player.HpLimit)
	if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterSpeedUp) {
		logMsg += fmt.Sprintf("【%s】`急速`触发 增加 %d 点速度 ", monster.Name, (monster.Speed*5)/100)
		monster.Speed += (monster.Speed * 5) / 100
		if (monster.Speed*5)/100 >= player.Speed {
			player.Hp = 0
			logMsg += fmt.Sprintf("【%s】的速度太快了，你已经跟不上了...", monster.Name)
		}
	}
	if player.Hp <= 0 {
		player.Hp = 0
		logMsg += fmt.Sprintf("玩家已死亡！\n")
		return logMsg
	}

	if monster.Speed > player.Speed*2 { // 速度超出两倍会触发多次伤害
		for i := 1; i <= int(monster.Speed/(player.Speed+1)); i++ {
			finalDmg = baseDmg
			if rand.Float64() <= monsterCritRate {
				finalDmg = int64(float64(baseDmg) * critMulti)
			}
			logMsg += fmt.Sprintf("【%s】触发追击对玩家造成 %d 点伤害 ", monster.Name, finalDmg)
			if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterRuthless) {
				finalDmg += (finalDmg * 30) / 100
				logMsg += fmt.Sprintf("`残暴`增伤 %d 点伤害 ", (finalDmg*30)/100)
			}
			if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterHuntBlood) {
				logMsg += fmt.Sprintf("【%s】`嗜血`增加了攻击 %d\n ", monster.Name, (monster.Attack*3)/100)
				monster.Attack += (monster.Attack * 3) / 100
			}
			if model.CheckHasSpecial(specials, model.SpecialsSolid) {
				logMsg += fmt.Sprintf("`坚固`抵挡了 %d 点伤害 ", (finalDmg*5)/100)
				finalDmg -= (finalDmg * 5) / 100
			}
			if model.CheckHasSpecial(specials, model.SpecialsSuperSolid) {
				logMsg += fmt.Sprintf("`超级坚固`抵挡了 %d 点伤害 ", (finalDmg*10)/100)
				finalDmg -= (finalDmg * 10) / 100
			}
			if finalDmg <= 0 {
				finalDmg = 1
			}
			if model.CheckHasSpecial(specials, model.SpecialsAggressive) {
				logMsg += fmt.Sprintf("`傲慢诅咒`增加了 %d 点收到的伤害 ", (finalDmg*15)/100)
				finalDmg += (finalDmg * 15) / 100
			}
			player.Hp -= finalDmg
			if player.Hp > 0 && model.CheckHasSpecial(specials, model.SpecialsStrong) {
				player.Hp += (player.HpLimit * 3) / 100
				logMsg += fmt.Sprintf("`强壮`恢复了 %d 点体魄 ", (player.HpLimit*3)/100)
			}
			if player.Hp > 0 && model.CheckHasSpecial(specials, model.SpecialsSuperStrong) {
				player.Hp += (player.HpLimit * 6) / 100
				logMsg += fmt.Sprintf("`超级强壮`恢复了 %d 点体魄 ", (player.HpLimit*6)/100)
			}
			if player.Hp < 0 {
				player.Hp = 0
			}
			logMsg += fmt.Sprintf(" 玩家剩余体魄：%d/%d\n", player.Hp, player.HpLimit)
			if model.CheckHasSpecial(monster.Special, model.SpecialsMonsterSpeedUp) {
				logMsg += fmt.Sprintf("【%s】`急速`触发 增加 %d 点速度 ", monster.Name, (monster.Speed*5)/100)
				monster.Speed += (monster.Speed * 5) / 100
				if (monster.Speed*5)/100 >= player.Speed {
					player.Hp = 0
					logMsg += fmt.Sprintf("【%s】的速度太快了，你已经跟不上了...", monster.Name)
				}
			}
			if player.Hp <= 0 {
				logMsg += fmt.Sprintf("玩家已死亡！\n")
				break
			}
		}
	}
	return logMsg
}
