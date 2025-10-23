package model

import "slices"

// 特效
const (
	SpecialsCritical       = "暴击"     // "暴击率提升 10%",
	SpecialsSuperCritical  = "超级暴击"   // "暴击率提升 20%",
	SpecialsMastery        = "精通"     // "暴击伤害提高 25%",
	SpecialsSuperMastery   = "超级精通"   // "暴击伤害提高 50%",
	SpecialsSpeedUp        = "轻盈"     // "闪避几率增加 10%",
	SpecialsSuperSpeedUp   = "超级轻盈"   // "闪避几率增加 20%",
	SpecialsSharp          = "尖锐"     // "攻击后额外造成当前攻击 5% 点伤害",
	SpecialsSuperSharp     = "超级尖锐"   // "攻击后额外造成当前攻击 10% 点伤害",
	SpecialsSolid          = "坚固"     // "单次伤害结算时减少 5% 受到的伤害",
	SpecialsSuperSolid     = "超级坚固"   // "单次伤害结算时减少 10% 受到的伤害",
	SpecialsStrong         = "强壮"     // "单次伤害结算后恢复 3% 体魄",
	SpecialsSuperStrong    = "超级强壮"   // "单次伤害结算后恢复 6% 体魄",
	SpecialsFast           = "迅捷"     // "战斗中每次出手后提高自身速度 2.5%",
	SpecialsSuperFast      = "超级迅捷"   // "战斗中每次出手后提高自身速度 5%",
	SpecialsSuckBlood      = "体魄偷取"   // "单次伤害结算后偷取造成伤害的 10%",
	SpecialsSuperSuckBlood = "超级体魄偷取" // "单次伤害结算后偷取造成伤害的 20%",
	//SpecialsSecondKill      = "秒杀"     // "单次伤害结算后有 2.5% 几率目标立刻死亡",
	//SpecialsSuperSecondKill = "超级秒杀"   // "单次伤害结算后有 5% 几率目标立刻死亡",

	// -------- 负面类 ------------

	SpecialsGreedy       = "贪婪诅咒" // "获取金币时减少收益的 20%",
	SpecialsWeak         = "脆弱诅咒" // "被暴击的几率提高 20%",
	SpecialsSlow         = "迟缓诅咒" // "有 30% 的几率最后出手",
	SpecialsAggressive   = "傲慢诅咒" // "单次受击伤害结算时增加 15% 受到的伤害",
	SpecialsAchillesHeel = "要害诅咒" // "单次受击存在 3% 即死几率",
	SpecialsBleed        = "流血诅咒" // "战斗结束后扣取体魄上限的 5%",
	SpecialsNoob         = "愚笨诅咒" // 失去闪避能力

	// ------- 秘境怪物特效 ------------

	SpecialsMonsterTough     = "强韧" // "单次伤害结算后恢复 10% 体魄",
	SpecialsMonsterRuthless  = "残暴" // "攻击后额外造成当前攻击 30% 点伤害",
	SpecialsMonsterHuntBlood = "嗜血" // "战斗中每次出手后会增加 3% 点攻击",
	SpecialsMonsterVengeful  = "复仇" // "死亡时造成一次无法抵挡的伤害，该伤害为攻击属性值的 70%",
	SpecialsMonsterVolatile  = "易爆" // "当前体魄低于 10% 时，会产生一次无法抵挡的当前体魄值的伤害，然后直接死亡",
	SpecialsMonsterEthereal  = "虚体" // "有 30% 的概率闪避攻击",
	SpecialsMonsterShield    = "屏障" // "单次受伤害结算后会增加 5% 点防御",
	SpecialsMonsterThorns    = "荆棘" // "收到伤害时反伤 10% ",
	SpecialsMonsterSpeedUp   = "急速" // "战斗中每次出手后提高自身速度 5%"
)

var SpecialsDescription = map[string]string{
	SpecialsCritical:       "暴击率提升 10%",
	SpecialsSuperCritical:  "暴击率提升 20%",
	SpecialsMastery:        "暴击伤害提高 25%",
	SpecialsSuperMastery:   "暴击伤害提高 50%",
	SpecialsSpeedUp:        "闪避几率增加 10%",
	SpecialsSuperSpeedUp:   "闪避几率增加 20%",
	SpecialsSharp:          "攻击后额外造成当前攻击 5% 点伤害",
	SpecialsSuperSharp:     "攻击后额外造成当前攻击 10% 点伤害",
	SpecialsSolid:          "单次伤害结算时减少 5% 受到的伤害",
	SpecialsSuperSolid:     "单次伤害结算时减少 10% 受到的伤害",
	SpecialsStrong:         "单次伤害结算后恢复 3% 体魄",
	SpecialsSuperStrong:    "单次伤害结算后恢复 6% 体魄",
	SpecialsFast:           "战斗中每次出手后提高自身速度 2.5%",
	SpecialsSuperFast:      "战斗中每次出手后提高自身速度 5%",
	SpecialsSuckBlood:      "单次伤害结算后偷取造成伤害的 10%",
	SpecialsSuperSuckBlood: "单次伤害结算后偷取造成伤害的 20%",
	//SpecialsSecondKill:      "单次伤害结算后有 2.5% 几率目标立刻死亡",
	//SpecialsSuperSecondKill: "单次伤害结算后有 5% 几率目标立刻死亡",

	// -------- 负面类 ------------

	SpecialsGreedy:       "获取金币时减少收益的 20%",
	SpecialsWeak:         "被暴击的几率提高 20%",
	SpecialsSlow:         "有 30% 的几率最后出手",
	SpecialsAggressive:   "单次受击伤害结算时增加 15% 受到的伤害",
	SpecialsAchillesHeel: "单次受击存在 3% 即死几率",
	SpecialsBleed:        "战斗结束后扣取体魄上限的 5%",
	SpecialsNoob:         "失去闪避能力",

	// ------- 秘境怪物特效 ------------

	SpecialsMonsterTough:     "单次伤害结算后恢复 10% 体魄",
	SpecialsMonsterRuthless:  "攻击后额外造成当前攻击 30% 点伤害",
	SpecialsMonsterHuntBlood: "战斗中每次出手后会增加 3% 点攻击",
	SpecialsMonsterVengeful:  "死亡时造成一次无法抵挡的伤害，该伤害为攻击属性值的 70%",
	SpecialsMonsterVolatile:  "当前体魄低于 10% 时，会产生一次无法抵挡的当前体魄值的伤害，然后直接死亡",
	SpecialsMonsterEthereal:  "有 30% 的概率闪避攻击",
	SpecialsMonsterShield:    "单次伤害结算后会增加 5% 点防御",
	SpecialsMonsterThorns:    "收到伤害时反伤 10% ",
	SpecialsMonsterSpeedUp:   "战斗中每次出手后提高自身速度 5%",
}

func CheckHasSpecial(specials []string, specialName string) bool {
	return slices.Contains(specials, specialName)
}
