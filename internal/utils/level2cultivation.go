package utils

import (
	"fmt"
	"strings"
)

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

// IntToRoman 将整数转换为罗马数字
// 参数:
//
//	num: 要转换的整数，范围应在1到3999之间
//
// 返回值:
//
//	string: 对应的罗马数字字符串，如果输入超出范围则返回"MAX"
func IntToRoman(num int64) string {
	// 检查输入是否在有效范围内
	if num < 1 || num > 3999 {
		return "MAX"
	}

	// 定义罗马数字对应的数值和符号
	var (
		values  = []int64{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	)

	var result strings.Builder

	// 从最大的数值开始转换
	for i := 0; i < len(values); i++ {
		// 当当前数值小于等于剩余数字时，添加对应的符号
		for num >= values[i] {
			result.WriteString(symbols[i])
			num -= values[i]
		}
	}

	return result.String()
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
