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
	EquipTypeHead      = iota // 头甲
	EquipTypeBody             // 胸甲
	EquipTypeArm              // 臂甲
	EquipTypeLeg              // 腿甲
	EquipTypeWeapon           // 武器
	EquipTypeAccessory        // 配饰
	EquipTypeMax
)

// 初始化随机数生成器
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomEquip 根据指定的等级范围和类型生成随机装备。
// 它根据给定的参数创建一个具有随机属性的装备，
// 包括名称、属性值和特殊效果。
//
// 参数:
//   - minLevel: 要生成的装备的最小等级
//   - maxLevel: 要生成的装备的最大等级
//   - equipType: 装备类型，-1 表示随机选择
//
// 返回值:
//   - *Equip: 指向新生成的装备结构体的指针
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
	attack := int64(float64(baseAttrs.attack) * (1.0 + float64(level)))
	defense := int64(float64(baseAttrs.defense) * (1.0 + float64(level)))
	speed := int64(float64(baseAttrs.speed) * (1.0 + float64(level)))
	hp := int64(float64(baseAttrs.hp) * (1.0 + float64(level)))

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
	case EquipTypeAccessory: // 配饰
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
			middles:  []string{"木", "石", "铁", "钢", "银", "金", "龙鳞"},
			suffixes: []string{"剑", "斧", "矛", "弓", "杖", "匕首", "锤"},
		}
	case int(EquipTypeAccessory): // 配饰
		return NameParts{
			prefixes: []string{"破旧的", "普通的", "精致的", "史诗的", "传奇的"},
			middles:  []string{"荧光", "星光", "辉光", "耀光", "银", "金", "龙鳞"},
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

// 可能的特效列表
var possibleSpecials = map[int][]string{
	int(EquipTypeHead): {
		SpecialsMastery, SpecialsSuperMastery, SpecialsSpeedUp, SpecialsSuperSpeedUp, SpecialsSolid, SpecialsSuperSolid, SpecialsStrong, SpecialsSuperStrong,

		SpecialsGreedy, SpecialsWeak, SpecialsSlow, SpecialsAggressive, SpecialsNoob, SpecialsAchillesHeel, SpecialsBleed,
	},
	int(EquipTypeBody): {
		SpecialsSpeedUp, SpecialsSuperSpeedUp, SpecialsSolid, SpecialsSuperSolid, SpecialsStrong, SpecialsSuperStrong,

		SpecialsGreedy, SpecialsWeak, SpecialsSlow, SpecialsAggressive, SpecialsNoob, SpecialsAchillesHeel, SpecialsBleed,
	},
	int(EquipTypeArm): {
		SpecialsSpeedUp, SpecialsSuperSpeedUp, SpecialsSolid, SpecialsSuperSolid, SpecialsStrong, SpecialsSuperStrong, SpecialsSuckBlood, SpecialsSuperSuckBlood,

		SpecialsGreedy, SpecialsWeak, SpecialsSlow, SpecialsAggressive, SpecialsNoob, SpecialsAchillesHeel, SpecialsBleed,
	},
	int(EquipTypeLeg): {
		SpecialsSpeedUp, SpecialsSuperSpeedUp, SpecialsSolid, SpecialsSuperSolid, SpecialsStrong, SpecialsSuperStrong, SpecialsFast, SpecialsSuperFast,

		SpecialsGreedy, SpecialsWeak, SpecialsSlow, SpecialsAggressive, SpecialsNoob, SpecialsAchillesHeel, SpecialsBleed,
	},
	int(EquipTypeWeapon): {
		SpecialsMastery, SpecialsSuperMastery, SpecialsCritical, SpecialsSuperCritical, SpecialsSharp, SpecialsSuperSharp, SpecialsSuckBlood, SpecialsSuperSuckBlood,

		SpecialsGreedy, SpecialsSlow, SpecialsNoob, SpecialsBleed,
	},
	int(EquipTypeAccessory): { // 配饰
		SpecialsMastery, SpecialsSuperMastery, SpecialsCritical, SpecialsSuperCritical, SpecialsSpeedUp, SpecialsSuperSpeedUp, SpecialsSharp, SpecialsSuperSharp, SpecialsSolid, SpecialsSuperSolid, SpecialsStrong, SpecialsSuperStrong, SpecialsFast, SpecialsSuperFast, SpecialsSuckBlood, SpecialsSuperSuckBlood,

		SpecialsGreedy, SpecialsWeak, SpecialsSlow, SpecialsAggressive, SpecialsNoob, SpecialsAchillesHeel, SpecialsBleed,
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
	case EquipTypeAccessory:
		typeName = "配饰"
	default:
		typeName = "装备"
	}

	var desc string
	desc = fmt.Sprintf("%s，一件%s，", e.Name, typeName)

	if len(e.Special) > 0 {
		desc += "拥有特殊效果："
		for i, s := range e.Special {
			if i > 0 {
				desc += "，"
			}
			desc += fmt.Sprintf("(%s) %s", s, SpecialsDescription[s])
		}
	} else {
		desc += "没有特殊效果。"
	}

	return desc
}

// UpgradeEquip 装备升级
func (e Equip) UpgradeEquip() Equip {
	e.Level = e.Level + 1
	e.Description = e.GenerateDescription()

	// 根据装备类型获取基础属性范围
	baseAttrs := getBaseAttributes(int(e.Type))
	e.Hp = e.Hp + int64(baseAttrs.hp)
	e.Attack = e.Attack + int64(baseAttrs.attack)
	e.Defense = e.Defense + int64(baseAttrs.defense)
	e.Speed = e.Speed + int64(baseAttrs.speed)

	return e
}
