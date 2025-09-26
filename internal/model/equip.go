package model

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Equip struct {
	Name        string
	Description string
	Type        uint  // 装备的部位
	Level       int64 // 锻造的等级
	Attack      int64
	Defense     int64
	Speed       int64
	Hp          int64
	Special     []string // 特效
}

const (
	EquipTypeHead   = iota // 头甲
	EquipTypeBody          // 胸甲
	EquipTypeArm           // 臂甲
	EquipTypeLeg           // 腿甲
	EquipTypeWeapon        // 武器
	EquipType              // 配饰
	EquipTypeMax
)

// 初始化随机数生成器
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomEquip 随机生成装备
func RandomEquip() *Equip {
	// 随机选择装备类型
	equipType := rand.Intn(EquipTypeMax)

	// 根据装备类型获取基础属性范围和可能的名称前缀/后缀
	baseAttrs := getBaseAttributes(equipType)
	nameParts := getEquipmentNameParts(equipType)

	// 随机生成装备名称
	name := generateEquipmentName(nameParts, equipType)

	var level int64
	level = 1

	// 根据等级和基础属性计算最终属性
	attack := int64(float64(baseAttrs.attack) * (1.0 + float64(level)*0.1))
	defense := int64(float64(baseAttrs.defense) * (1.0 + float64(level)*0.1))
	speed := int64(float64(baseAttrs.speed) * (1.0 + float64(level)*0.1))
	hp := int64(float64(baseAttrs.hp) * (1.0 + float64(level)*0.1))

	// 随机生成特效
	//specials := generateSpecials(equipType, level)
	var specials []string

	// 生成描述
	description := generateDescription(equipType, level, specials)

	return &Equip{
		Name:        name,
		Description: description,
		Type:        uint(equipType),
		Level:       level,
		Attack:      attack,
		Defense:     defense,
		Speed:       speed,
		Hp:          hp,
		Special:     specials,
	}
}

// 基础属性结构
type baseAttributes struct {
	attack  int
	defense int
	speed   int
	hp      int
}

// 根据装备类型获取基础属性范围
func getBaseAttributes(equipType int) baseAttributes {
	switch equipType {
	case EquipTypeHead:
		return baseAttributes{0, 1, 0, 5}
	case EquipTypeBody:
		return baseAttributes{0, 2, 0, 0}
	case EquipTypeArm:
		return baseAttributes{1, 1, 0, 0}
	case EquipTypeLeg:
		return baseAttributes{0, 1, 1, 0}
	case EquipTypeWeapon:
		return baseAttributes{2, 0, 0, 0}
	case EquipType: // 配饰
		return baseAttributes{0, 0, 1, 5}
	default:
		return baseAttributes{0, 0, 0, 0}
	}
}

// NameParts 装备名称部分
type NameParts struct {
	prefixes []string
	middles  []string
	suffixes []string
}

// 根据装备类型获取名称部分
func getEquipmentNameParts(equipType int) NameParts {
	switch equipType {
	case int(EquipTypeHead):
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"布", "皮", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"头盔", "帽子", "头冠", "面罩", "头饰"},
		}
	case int(EquipTypeBody):
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"布", "皮", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"胸甲", "铠甲", "外套", "长袍", "背心"},
		}
	case int(EquipTypeArm):
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"布", "皮", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"臂甲", "护腕", "手套", "护手", "护臂"},
		}
	case int(EquipTypeLeg):
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"布", "皮", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"腿甲", "护腿", "战靴", "长靴", "胫甲"},
		}
	case int(EquipTypeWeapon):
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"布", "皮", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"剑", "斧", "矛", "弓", "杖", "匕首", "锤"},
		}
	case int(EquipType): // 配饰
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"布", "皮", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"项链", "戒指", "手镯", "徽章", "护身符"},
		}
	default:
		return NameParts{
			prefixes: []string{"普通的"},
			middles:  []string{""},
			suffixes: []string{"装备"},
		}
	}
}

// 生成装备名称
func generateEquipmentName(parts NameParts, equipType int) string {
	prefix := parts.prefixes[rand.Intn(len(parts.prefixes))]
	middle := parts.middles[rand.Intn(len(parts.middles))]
	suffix := parts.suffixes[rand.Intn(len(parts.suffixes))]

	return prefix + middle + suffix
}

// 可能的特效列表
var possibleSpecials = map[int][]string{
	int(EquipTypeHead): {
		"增加5%暴击率",
		"免疫眩晕",
		"提高10%魔法抗性",
		"增加视野范围",
		"减少受到的远程伤害",
	},
	int(EquipTypeBody): {
		"增加10%最大生命值",
		"每秒钟恢复1%生命值",
		"减少10%受到的伤害",
		"增加背包容量",
		"提高所有属性2%",
	},
	int(EquipTypeArm): {
		"增加10%攻击速度",
		"提高5%命中率",
		"减少技能冷却时间",
		"增加10%暴击伤害",
		"有几率造成额外伤害",
	},
	int(EquipTypeLeg): {
		"增加15%移动速度",
		"减少受到的移动限制效果",
		"提高闪避率",
		"增加跳跃高度",
		"免疫减速",
	},
	int(EquipTypeWeapon): {
		"增加10%攻击力",
		"有几率造成双倍伤害",
		"攻击时吸取生命值",
		"对特定类型敌人造成额外伤害",
		"攻击时有几率击晕敌人",
	},
	int(EquipType): { // 配饰
		"增加所有属性5%",
		"提高经验获取率",
		"增加金币掉落率",
		"提高元素抗性",
		"死后有几率复活一次",
	},
}

// 生成装备特效
func generateSpecials(equipType int, level int64) []string {
	var specials []string
	maxSpecials := 1

	// 等级越高，可能的特效越多
	if level >= 5 {
		maxSpecials = 2
	}
	if level >= 8 {
		maxSpecials = 3
	}

	// 随机决定实际特效数量
	numSpecials := rand.Intn(maxSpecials) + 1

	// 从可能的特效中随机选择
	possible := possibleSpecials[equipType]
	selected := make(map[int]bool)

	for i := 0; i < numSpecials; i++ {
		// 确保不重复选择同一个特效
		idx := rand.Intn(len(possible))
		for selected[idx] {
			idx = rand.Intn(len(possible))
		}
		selected[idx] = true
		specials = append(specials, possible[idx])
	}

	return specials
}

// 生成装备描述
func generateDescription(equipType int, level int64, specials []string) string {
	var typeName string
	switch equipType {
	case int(EquipTypeHead):
		typeName = "头甲"
	case int(EquipTypeBody):
		typeName = "胸甲"
	case int(EquipTypeArm):
		typeName = "臂甲"
	case int(EquipTypeLeg):
		typeName = "腿甲"
	case int(EquipTypeWeapon):
		typeName = "武器"
	case int(EquipType):
		typeName = "配饰"
	default:
		typeName = "装备"
	}

	var desc strings.Builder
	fmt.Sprintf("一件%d级的%s，", level, typeName)

	if len(specials) > 0 {
		desc.WriteString("拥有特殊效果：")
		for i, s := range specials {
			if i > 0 {
				desc.WriteString("，")
			}
			desc.WriteString(s)
		}
	} else {
		desc.WriteString("没有特殊效果。")
	}

	return desc.String()
}
