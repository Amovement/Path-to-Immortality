package utils

import "fmt"

// 定义境界名称列表，顺序从低到高
var realms = []string{
	"引气境",
	"筑基境",
	"金丹境",
	"元婴境",
	"化神境",
	"炼虚境",
	"合体境",
	"大乘境",
	"渡劫境",
}

// GetCultivationByLevel 根据等级获取对应的修仙境界名称
func GetCultivationByLevel(level int) string {
	// 特殊处理初始状态
	if level == 0 {
		return "灵智未启"
	}

	// 调整等级基数（从1开始计算）
	adjustedLevel := level - 1

	// 渡劫境特殊处理（最后一个境界）
	lastRealmIndex := len(realms) - 1
	tribulationRealm := realms[lastRealmIndex]

	// 计算每个常规境界的等级范围（每个境界30级：3个阶段×10级）
	normalRealmCount := lastRealmIndex
	normalTotalLevels := normalRealmCount * 30

	// 如果超过常规境界总等级，进入渡劫境
	if adjustedLevel >= normalTotalLevels {
		tribulationLevel := adjustedLevel - normalTotalLevels + 1
		if tribulationLevel >= 10 {
			return fmt.Sprintf("%s圆满（飞升）", tribulationRealm)
		}
		return fmt.Sprintf("%s%d重", tribulationRealm, tribulationLevel)
	}

	// 计算常规境界信息
	realmIndex := adjustedLevel / 30
	realm := realms[realmIndex]
	stageLevel := adjustedLevel % 30    // 境界内的等级（0-29）
	stage := stageLevel / 10            // 阶段（0-2：前期/中期/后期）
	stageOrder := (stageLevel % 10) + 1 // 阶段内的阶数（1-9）

	// 阶段名称映射
	stageNames := []string{"前期", "中期", "后期"}

	// 处理突破等级（每个阶段的第10级）
	if stageOrder == 10 {
		return fmt.Sprintf("%s%s巅峰", realm, stageNames[stage])
	}

	// 常规等级
	return fmt.Sprintf("%s%s%d阶", realm, stageNames[stage], stageOrder)
}
