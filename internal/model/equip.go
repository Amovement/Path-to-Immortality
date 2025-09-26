package model

import (
	"fmt"
	"math/rand"
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
func RandomEquip(minLevel, maxLevel int64, equipType int) *Equip {
	level := rand.Int63n(maxLevel-minLevel+1) + minLevel

	// 随机选择装备类型
	if equipType == -1 {
		equipType = rand.Intn(EquipTypeMax)
	}

	// 根据装备类型获取基础属性范围和可能的名称前缀/后缀
	baseAttrs := getBaseAttributes(equipType)
	nameParts := getEquipmentNameParts(equipType)

	// 随机生成装备名称
	name := generateEquipmentName(nameParts, level)

	// 根据等级和基础属性计算最终属性
	attack := int64(float64(baseAttrs.attack) * (1.0 + float64(level)*0.5))
	defense := int64(float64(baseAttrs.defense) * (1.0 + float64(level)*0.5))
	speed := int64(float64(baseAttrs.speed) * (1.0 + float64(level)*0.5))
	hp := int64(float64(baseAttrs.hp) * (1.0 + float64(level)*0.5))

	// 随机生成特效
	specials := generateSpecials(equipType, level)

	equipNew := Equip{
		Name:    name,
		Type:    uint(equipType),
		Level:   level,
		Attack:  attack,
		Defense: defense,
		Speed:   speed,
		Hp:      hp,
		Special: specials,
	}

	// 生成描述
	equipNew.Description = equipNew.GenerateDescription()

	return &equipNew
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

// generateEquipmentName 根据给定的名称部件和等级生成装备名称
// 参数:
//   - parts: NameParts类型，包含前缀、中间部分和后缀的字符串切片
//   - level: int64类型，表示装备等级，用于决定前缀的选择
//
// 返回值:
//   - string: 生成的装备名称字符串
func generateEquipmentName(parts NameParts, level int64) string {
	var prefix string
	// 根据等级选择前缀
	if level <= 10 {
		prefix = parts.prefixes[0]
	} else if level <= 20 {
		prefix = parts.prefixes[1]
	} else if level <= 30 {
		prefix = parts.prefixes[2]
	} else if level <= 40 {
		prefix = parts.prefixes[3]
	} else {
		prefix = parts.prefixes[4]
	}

	// 随机选择中间部分和后缀
	middle := parts.middles[rand.Intn(len(parts.middles))]
	suffix := parts.suffixes[rand.Intn(len(parts.suffixes))]

	// 组合前缀、中间部分和后缀生成完整名称
	return prefix + middle + suffix
}

const (
	SpecialsCritical = "暴击"
	SpecialsSolid    = "坚固"
)

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
	if level > 10 {
		maxSpecials = 2
	}
	if level > 20 {
		maxSpecials = 3
	}
	if level > 30 {
		maxSpecials = 4
	}

	// 随机决定实际特效数量
	numSpecials := rand.Intn(maxSpecials + 1)

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

// GenerateDescription 生成装备描述
func (e Equip) GenerateDescription() string {
	var typeName string
	switch e.Type {
	case EquipTypeHead:
		typeName = "头甲"
	case EquipTypeBody:
		typeName = "胸甲"
	case EquipTypeArm:
		typeName = "臂甲"
	case EquipTypeLeg:
		typeName = "腿甲"
	case EquipTypeWeapon:
		typeName = "武器"
	case EquipType:
		typeName = "配饰"
	default:
		typeName = "装备"
	}

	var desc string
	desc = fmt.Sprintf("一件%d级的%s，", e.Level, typeName)

	if len(e.Special) > 0 {
		desc += "拥有特殊效果："
		for i, s := range e.Special {
			if i > 0 {
				desc += "，"
			}
			desc += s
		}
	} else {
		desc += "没有特殊效果。"
	}

	return desc
}
