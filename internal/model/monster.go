package model

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"math/rand"
	"time"
)

type Monster struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Hp          int64    `json:"hp"`
	HpLimit     int64    `json:"hpLimit"`
	Attack      int64    `json:"attack"`
	Defense     int64    `json:"defense"`
	Speed       int64    `json:"speed"`
	Level       int64    `json:"level"`
	Special     []string `json:"special"`
	Cultivation string   `json:"cultivation"`
}

// 怪物名称前缀和后缀，用于生成随机名称
var (
	chineseTime     = []string{"子时", "丑时", "寅时", "卯时", "辰时", "巳时", "午时", "未时", "申时", "酉时", "戌时", "亥时"}
	monsterPrefixes = []string{"狂暴", "暗影", "烈焰", "寒冰", "剧毒", "钢铁", "雷霆", "大地", "风暴", "虚空"}
	monsterSuffixes = []string{"贪狼", "极北龙", "凤凰", "白虎", "泰坦", "圣狮", "蝎子精", "青龙", "九尾狐", "独角兽", "战鹰", "巨像", "幻狐", "蛟龙", "多头蛇", "灵鹤", "守卫"}
	specialSkills   = []string{
		SpecialsMonsterTough,
		SpecialsMonsterRuthless,
		SpecialsMonsterHuntBlood,
		SpecialsMonsterVengeful,
		SpecialsMonsterVolatile,
		SpecialsMonsterEthereal,
		SpecialsMonsterShield,
		SpecialsMonsterThorns,
	}
)

// GenerateRandomMythicMonster 根据等级生成随机怪物
//
//	level: 怪物等级
//	返回值: 生成的随机怪物
func GenerateRandomMythicMonster(level int64) *Monster {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 基础属性 = 等级 * 随机系数 + 随机基础值
	// 确保属性随等级提升而增长，同时保持一定随机性
	hpLimit := int64(rand.Intn(50)+5) + level*int64(rand.Intn(25)+15)
	attack := int64(rand.Intn(10)+1) + level*int64(rand.Intn(5)+3)
	defense := int64(rand.Intn(10)+1) + level*int64(rand.Intn(5)+3)
	speed := int64(rand.Intn(10)+1) + level*int64(rand.Intn(5)+3)

	// 随机生成名称
	name := chineseTime[rand.Intn(len(chineseTime))] +
		monsterPrefixes[rand.Intn(len(monsterPrefixes))] +
		monsterSuffixes[rand.Intn(len(monsterSuffixes))]

	// 随机生成特殊技能（0-3个）
	specials := make([]string, 0)
	// 等级越高技能数越多
	skillCount := rand.Intn(int(level/10)+2) + int(level/10)
	skillCount = int(min(int64(skillCount), int64(len(specialSkills))))
	for i := 0; i < skillCount; i++ {
		skill := specialSkills[rand.Intn(len(specialSkills))]
		// 避免重复技能
		if !CheckHasSpecial(specials, skill) {
			specials = append(specials, skill)
		}
	}

	return &Monster{
		ID:          uint(rand.Uint64() % 1000), // 生成一个随机ID
		Name:        name,
		Hp:          hpLimit, // 初始生命值为上限值
		HpLimit:     hpLimit,
		Attack:      attack,
		Defense:     defense,
		Speed:       speed,
		Level:       level,
		Special:     specials,
		Cultivation: utils.GetCultivationByLevel(int(level)),
	}
}
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
