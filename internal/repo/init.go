package repo

import "github.com/Amovement/Path-to-Immortality-WASM/internal/model"

var monsters = []model.Monster{
	// 引气境怪物（1-30级）- 低阶生灵，属性温和，适合新手过渡
	{ID: 1, Name: "赤毛灵狐", Hp: 50, HpLimit: 50, Attack: 5, Defense: 1, Speed: 3, Level: 1},
	{ID: 2, Name: "铁皮妖鼠", Hp: 80, HpLimit: 80, Attack: 7, Defense: 3, Speed: 2, Level: 3},
	{ID: 3, Name: "腐气僵尸", Hp: 150, HpLimit: 150, Attack: 12, Defense: 5, Speed: 4, Level: 5},
	{ID: 4, Name: "青面山魈", Hp: 250, HpLimit: 250, Attack: 20, Defense: 8, Speed: 6, Level: 8},
	{ID: 5, Name: "启灵魔仆", Hp: 400, HpLimit: 400, Attack: 30, Defense: 12, Speed: 9, Level: 12},
	{ID: 6, Name: "百年石灵", Hp: 600, HpLimit: 600, Attack: 25, Defense: 20, Speed: 3, Level: 15},
	{ID: 7, Name: "风系小妖", Hp: 350, HpLimit: 350, Attack: 35, Defense: 10, Speed: 18, Level: 18},
	{ID: 8, Name: "无名散修", Hp: 500, HpLimit: 500, Attack: 40, Defense: 15, Speed: 12, Level: 25},

	// 筑基境怪物（31-60 级）- 属性提升显著，加入元素特性，增加策略性
	{ID: 9, Name: "筑基境傀儡", Hp: 1200, HpLimit: 1200, Attack: 60, Defense: 30, Speed: 15, Level: 32},
	{ID: 10, Name: "火鸦妖王", Hp: 1500, HpLimit: 1500, Attack: 75, Defense: 25, Speed: 20, Level: 35},
	{ID: 11, Name: "玄冰精怪", Hp: 1800, HpLimit: 1800, Attack: 55, Defense: 40, Speed: 10, Level: 40},
	{ID: 12, Name: "金丹长老分身", Hp: 2200, HpLimit: 2200, Attack: 85, Defense: 35, Speed: 22, Level: 45},
	{ID: 13, Name: "血炼魔徒", Hp: 2800, HpLimit: 2800, Attack: 95, Defense: 45, Speed: 18, Level: 50},
	{ID: 14, Name: "雷纹巨猿", Hp: 3500, HpLimit: 3500, Attack: 110, Defense: 50, Speed: 15, Level: 55},

	// 金丹境怪物（61-90 级）- 融入 "灵智" 特性，属性分化明显
	{ID: 15, Name: "护法", Hp: 5000, HpLimit: 5000, Attack: 150, Defense: 70, Speed: 25, Level: 65},
	{ID: 16, Name: "骨甲尸王", Hp: 6500, HpLimit: 6500, Attack: 130, Defense: 90, Speed: 18, Level: 70},
	{ID: 17, Name: "风灵圣女", Hp: 4800, HpLimit: 4800, Attack: 170, Defense: 60, Speed: 35, Level: 75},
	{ID: 18, Name: "宗门长老", Hp: 8000, HpLimit: 8000, Attack: 190, Defense: 85, Speed: 30, Level: 80},
	{ID: 19, Name: "熔岩古龙", Hp: 10000, HpLimit: 10000, Attack: 210, Defense: 100, Speed: 22, Level: 85},

	// 元婴境怪物（91-120 级）- 加入 "法则" 元素，属性跳跃式提升
	{ID: 20, Name: "上古傀儡", Hp: 15000, HpLimit: 15000, Attack: 280, Defense: 150, Speed: 30, Level: 95},
	{ID: 21, Name: "玄天蛇姬", Hp: 18000, HpLimit: 18000, Attack: 320, Defense: 130, Speed: 40, Level: 100},
	{ID: 22, Name: "血河老怪", Hp: 22000, HpLimit: 22000, Attack: 350, Defense: 160, Speed: 35, Level: 105},
	{ID: 23, Name: "雷劫剑灵", Hp: 16000, HpLimit: 16000, Attack: 400, Defense: 120, Speed: 50, Level: 110},
	{ID: 24, Name: "元婴境大长老", Hp: 25000, HpLimit: 25000, Attack: 380, Defense: 180, Speed: 45, Level: 115},

	// 化神境怪物（121-150 级）- 侧重 "空间 / 时间" 特性，属性偏向极端
	{ID: 25, Name: "化神妖兽", Hp: 35000, HpLimit: 35000, Attack: 480, Defense: 250, Speed: 40, Level: 125},
	{ID: 26, Name: "空间裂缝兽", Hp: 30000, HpLimit: 30000, Attack: 520, Defense: 220, Speed: 55, Level: 130},
	{ID: 27, Name: "上古战魂", Hp: 45000, HpLimit: 45000, Attack: 550, Defense: 280, Speed: 45, Level: 135},
	{ID: 28, Name: "炼虚长老分身", Hp: 50000, HpLimit: 50000, Attack: 600, Defense: 300, Speed: 50, Level: 140},
	{ID: 29, Name: "混沌石灵", Hp: 60000, HpLimit: 60000, Attack: 500, Defense: 400, Speed: 30, Level: 145},

	// 炼虚境怪物（151-180 级）- 融入 "天地灵气" 设定，属性全面强化
	{ID: 30, Name: "魔将", Hp: 80000, HpLimit: 80000, Attack: 750, Defense: 450, Speed: 55, Level: 155},
	{ID: 31, Name: "九天玄女", Hp: 70000, HpLimit: 70000, Attack: 800, Defense: 400, Speed: 65, Level: 160},
	{ID: 32, Name: "万尸之主", Hp: 100000, HpLimit: 100000, Attack: 700, Defense: 500, Speed: 50, Level: 165},
	{ID: 33, Name: "宗主化身", Hp: 90000, HpLimit: 90000, Attack: 850, Defense: 480, Speed: 60, Level: 170},
	{ID: 34, Name: "洪荒巨兽", Hp: 120000, HpLimit: 120000, Attack: 900, Defense: 550, Speed: 45, Level: 175},

	// 合体境怪物（181-210 级）- 加入 "本源" 特性，属性接近神祇
	{ID: 35, Name: "合体仙将", Hp: 150000, HpLimit: 150000, Attack: 1100, Defense: 650, Speed: 70, Level: 185},
	{ID: 36, Name: "灭世魔魂", Hp: 180000, HpLimit: 180000, Attack: 1200, Defense: 600, Speed: 75, Level: 190},
	{ID: 37, Name: "天道守护者", Hp: 200000, HpLimit: 200000, Attack: 1000, Defense: 800, Speed: 65, Level: 195},
	{ID: 38, Name: "圣祖", Hp: 220000, HpLimit: 220000, Attack: 1300, Defense: 750, Speed: 80, Level: 200},
	{ID: 39, Name: "混沌之影", Hp: 250000, HpLimit: 250000, Attack: 1400, Defense: 700, Speed: 90, Level: 205},

	// 大乘境、渡劫境怪物（211-250 级）- 匹配 "天劫" 设定，属性达到凡界顶峰
	{ID: 40, Name: "雷劫之灵", Hp: 300000, HpLimit: 300000, Attack: 1600, Defense: 800, Speed: 100, Level: 215},
	{ID: 41, Name: "空间法则兽", Hp: 280000, HpLimit: 280000, Attack: 1700, Defense: 750, Speed: 110, Level: 220},
	{ID: 42, Name: "上古仙尊残魂", Hp: 350000, HpLimit: 350000, Attack: 1800, Defense: 900, Speed: 95, Level: 225},
	{ID: 43, Name: "渡劫天魔", Hp: 400000, HpLimit: 400000, Attack: 2000, Defense: 850, Speed: 105, Level: 230},
	{ID: 44, Name: "天道执法者", Hp: 500000, HpLimit: 500000, Attack: 2200, Defense: 1000, Speed: 115, Level: 240},
	{ID: 45, Name: "鸿蒙本源兽", Hp: 600000, HpLimit: 600000, Attack: 2500, Defense: 1200, Speed: 120, Level: 250},
}

// Challenges 50个挑战任务（每5级一个梯度，覆盖1-250级）
var Challenges = []model.Challenge{
	{ID: 1, LevelLimit: 30, Title: "新手林地试炼", Gold: 1},   // 1级
	{ID: 2, LevelLimit: 30, Title: "乱葬岗清剿", Gold: 3},    // 5级
	{ID: 3, LevelLimit: 30, Title: "黑风洞探秘", Gold: 5},    // 10级
	{ID: 4, LevelLimit: 30, Title: "灵草园守护", Gold: 8},    // 15级
	{ID: 5, LevelLimit: 30, Title: "引气期修士考验", Gold: 12}, // 20级
	{ID: 6, LevelLimit: 30, Title: "妖兽潮拦截", Gold: 20},   // 25级
	{ID: 7, LevelLimit: 30, Title: "引气境圆满挑战", Gold: 30}, // 30级

	{ID: 8, LevelLimit: 60, Title: "筑基秘境初探", Gold: 50},    // 35级
	{ID: 9, LevelLimit: 60, Title: "火鸦巢清理", Gold: 60},     // 40级
	{ID: 10, LevelLimit: 60, Title: "玄冰窟探险", Gold: 70},    // 45级
	{ID: 11, LevelLimit: 60, Title: "筑基期入门考核", Gold: 80},  // 50级
	{ID: 12, LevelLimit: 60, Title: "血魔谷围剿", Gold: 90},    // 55级
	{ID: 13, LevelLimit: 60, Title: "筑基境圆满挑战", Gold: 100}, // 60级

	{ID: 14, LevelLimit: 90, Title: "金丹洞府开启", Gold: 120},  // 65级
	{ID: 15, LevelLimit: 90, Title: "尸王殿破局", Gold: 140},   // 70级
	{ID: 16, LevelLimit: 90, Title: "风灵谷试炼", Gold: 160},   // 75级
	{ID: 17, LevelLimit: 90, Title: "金丹长老考验", Gold: 180},  // 80级
	{ID: 18, LevelLimit: 90, Title: "古龙巢穴探秘", Gold: 200},  // 85级
	{ID: 19, LevelLimit: 90, Title: "金丹境圆满挑战", Gold: 220}, // 90级

	{ID: 20, LevelLimit: 120, Title: "元婴遗迹探索", Gold: 250},  // 95级
	{ID: 21, LevelLimit: 120, Title: "玄天蛇窟历险", Gold: 280},  // 100级
	{ID: 22, LevelLimit: 120, Title: "血河源头镇压", Gold: 310},  // 105级
	{ID: 23, LevelLimit: 120, Title: "雷劫剑灵降服", Gold: 340},  // 110级
	{ID: 24, LevelLimit: 120, Title: "元婴大典挑战", Gold: 370},  // 115级
	{ID: 25, LevelLimit: 120, Title: "元婴境圆满挑战", Gold: 400}, // 120级

	{ID: 26, LevelLimit: 150, Title: "化神秘境闯关", Gold: 450},  // 125级
	{ID: 27, LevelLimit: 150, Title: "空间裂缝平定", Gold: 500},  // 130级
	{ID: 28, LevelLimit: 150, Title: "上古战场历练", Gold: 550},  // 135级
	{ID: 29, LevelLimit: 150, Title: "化神长老对决", Gold: 600},  // 140级
	{ID: 30, LevelLimit: 150, Title: "混沌石林试炼", Gold: 650},  // 145级
	{ID: 31, LevelLimit: 150, Title: "化神境圆满挑战", Gold: 700}, // 150级

	{ID: 32, LevelLimit: 180, Title: "魔窟征伐", Gold: 770},     // 155级
	{ID: 33, LevelLimit: 180, Title: "九天瑶池试炼", Gold: 800},   // 160级
	{ID: 34, LevelLimit: 180, Title: "万尸窟净化", Gold: 870},    // 165级
	{ID: 35, LevelLimit: 180, Title: "宗主宝座争夺", Gold: 900},   // 170级
	{ID: 36, LevelLimit: 180, Title: "洪荒兽域探险", Gold: 970},   // 175级
	{ID: 37, LevelLimit: 180, Title: "炼虚境圆满挑战", Gold: 1000}, // 180级

	{ID: 38, LevelLimit: 210, Title: "仙域开启", Gold: 1100},    // 185级
	{ID: 39, LevelLimit: 210, Title: "灭世魔魂封印", Gold: 1200},  // 190级
	{ID: 40, LevelLimit: 210, Title: "天道试炼场", Gold: 1300},   // 195级
	{ID: 41, LevelLimit: 210, Title: "圣祖传承考验", Gold: 1400},  // 200级
	{ID: 42, LevelLimit: 210, Title: "混沌空间历练", Gold: 1500},  // 205级
	{ID: 43, LevelLimit: 210, Title: "合体境圆满挑战", Gold: 1600}, // 210级

	{ID: 44, LevelLimit: 240, Title: "一重雷劫挑战", Gold: 1800},  // 215级
	{ID: 45, LevelLimit: 240, Title: "空间法则试炼", Gold: 2000},  // 220级
	{ID: 46, LevelLimit: 240, Title: "仙尊传承争夺", Gold: 3000},  // 225级
	{ID: 47, LevelLimit: 240, Title: "天魔巢穴围剿", Gold: 4000},  // 230级
	{ID: 48, LevelLimit: 240, Title: "天道法则领悟", Gold: 5000},  // 235级
	{ID: 49, LevelLimit: 240, Title: "执法者考验", Gold: 6000},   // 240级
	{ID: 50, LevelLimit: 240, Title: "鸿蒙本源试炼", Gold: 7000},  // 245级
	{ID: 51, LevelLimit: 240, Title: "渡劫境圆满挑战", Gold: 8000}, // 250级
}

// 挑战-怪物关联数据（按每级3点属性成长平衡难度）
var challengeMonsters = []model.ChallengeMonster{
	// 引气境挑战（1-30级）
	{ChallengeID: 1, MonsterID: 1, Count: 3}, {ChallengeID: 1, MonsterID: 2, Count: 1},
	{ChallengeID: 2, MonsterID: 3, Count: 2}, {ChallengeID: 2, MonsterID: 4, Count: 1},
	{ChallengeID: 3, MonsterID: 4, Count: 3}, {ChallengeID: 3, MonsterID: 5, Count: 1},
	{ChallengeID: 4, MonsterID: 5, Count: 2}, {ChallengeID: 4, MonsterID: 7, Count: 2},
	{ChallengeID: 5, MonsterID: 7, Count: 3}, {ChallengeID: 5, MonsterID: 8, Count: 1},
	{ChallengeID: 6, MonsterID: 4, Count: 5}, {ChallengeID: 6, MonsterID: 5, Count: 3}, {ChallengeID: 6, MonsterID: 8, Count: 2},
	{ChallengeID: 7, MonsterID: 8, Count: 5}, {ChallengeID: 7, MonsterID: 6, Count: 3},

	// 金丹境挑战（31-60级）
	{ChallengeID: 8, MonsterID: 9, Count: 2}, {ChallengeID: 8, MonsterID: 10, Count: 1},
	{ChallengeID: 9, MonsterID: 10, Count: 3}, {ChallengeID: 9, MonsterID: 11, Count: 1},
	{ChallengeID: 10, MonsterID: 11, Count: 2}, {ChallengeID: 10, MonsterID: 12, Count: 1},
	{ChallengeID: 11, MonsterID: 12, Count: 2}, {ChallengeID: 11, MonsterID: 13, Count: 1},
	{ChallengeID: 12, MonsterID: 13, Count: 3}, {ChallengeID: 12, MonsterID: 14, Count: 1},
	{ChallengeID: 13, MonsterID: 14, Count: 2}, {ChallengeID: 13, MonsterID: 12, Count: 2},

	// 元婴境挑战（61-90级）
	{ChallengeID: 14, MonsterID: 15, Count: 2}, {ChallengeID: 14, MonsterID: 16, Count: 1},
	{ChallengeID: 15, MonsterID: 16, Count: 2}, {ChallengeID: 15, MonsterID: 17, Count: 1},
	{ChallengeID: 16, MonsterID: 17, Count: 2}, {ChallengeID: 16, MonsterID: 18, Count: 1},
	{ChallengeID: 17, MonsterID: 18, Count: 2}, {ChallengeID: 17, MonsterID: 19, Count: 1},
	{ChallengeID: 18, MonsterID: 19, Count: 2}, {ChallengeID: 18, MonsterID: 18, Count: 1},
	{ChallengeID: 19, MonsterID: 19, Count: 3}, {ChallengeID: 19, MonsterID: 17, Count: 2},

	// 化神境挑战（91-120级）
	{ChallengeID: 20, MonsterID: 20, Count: 2}, {ChallengeID: 20, MonsterID: 21, Count: 1},
	{ChallengeID: 21, MonsterID: 21, Count: 2}, {ChallengeID: 21, MonsterID: 22, Count: 1},
	{ChallengeID: 22, MonsterID: 22, Count: 2}, {ChallengeID: 22, MonsterID: 23, Count: 1},
	{ChallengeID: 23, MonsterID: 23, Count: 2}, {ChallengeID: 23, MonsterID: 24, Count: 1},
	{ChallengeID: 24, MonsterID: 24, Count: 2}, {ChallengeID: 24, MonsterID: 22, Count: 1},
	{ChallengeID: 25, MonsterID: 24, Count: 3}, {ChallengeID: 25, MonsterID: 23, Count: 2},

	// 炼虚境挑战（121-150级）
	{ChallengeID: 26, MonsterID: 25, Count: 2}, {ChallengeID: 26, MonsterID: 26, Count: 1},
	{ChallengeID: 27, MonsterID: 26, Count: 2}, {ChallengeID: 27, MonsterID: 27, Count: 1},
	{ChallengeID: 28, MonsterID: 27, Count: 2}, {ChallengeID: 28, MonsterID: 28, Count: 1},
	{ChallengeID: 29, MonsterID: 28, Count: 2}, {ChallengeID: 29, MonsterID: 29, Count: 1},
	{ChallengeID: 30, MonsterID: 29, Count: 2}, {ChallengeID: 30, MonsterID: 27, Count: 1},
	{ChallengeID: 31, MonsterID: 29, Count: 3}, {ChallengeID: 31, MonsterID: 28, Count: 2},

	// 合体境挑战（151-180级）
	{ChallengeID: 32, MonsterID: 30, Count: 2}, {ChallengeID: 32, MonsterID: 31, Count: 1},
	{ChallengeID: 33, MonsterID: 31, Count: 2}, {ChallengeID: 33, MonsterID: 32, Count: 1},
	{ChallengeID: 34, MonsterID: 32, Count: 2}, {ChallengeID: 34, MonsterID: 33, Count: 1},
	{ChallengeID: 35, MonsterID: 33, Count: 2}, {ChallengeID: 35, MonsterID: 34, Count: 1},
	{ChallengeID: 36, MonsterID: 34, Count: 2}, {ChallengeID: 36, MonsterID: 32, Count: 1},
	{ChallengeID: 37, MonsterID: 34, Count: 3}, {ChallengeID: 37, MonsterID: 33, Count: 2},

	// 大乘境挑战（181-210级）
	{ChallengeID: 38, MonsterID: 35, Count: 2}, {ChallengeID: 38, MonsterID: 36, Count: 1},
	{ChallengeID: 39, MonsterID: 36, Count: 2}, {ChallengeID: 39, MonsterID: 37, Count: 1},
	{ChallengeID: 40, MonsterID: 37, Count: 2}, {ChallengeID: 40, MonsterID: 38, Count: 1},
	{ChallengeID: 41, MonsterID: 38, Count: 2}, {ChallengeID: 41, MonsterID: 39, Count: 1},
	{ChallengeID: 42, MonsterID: 39, Count: 2}, {ChallengeID: 42, MonsterID: 37, Count: 1},
	{ChallengeID: 43, MonsterID: 39, Count: 3}, {ChallengeID: 43, MonsterID: 38, Count: 2},

	// 渡劫境挑战（211-250级）
	{ChallengeID: 44, MonsterID: 40, Count: 1}, {ChallengeID: 44, MonsterID: 23, Count: 3},
	{ChallengeID: 45, MonsterID: 41, Count: 1}, {ChallengeID: 45, MonsterID: 26, Count: 3},
	{ChallengeID: 46, MonsterID: 42, Count: 1}, {ChallengeID: 46, MonsterID: 28, Count: 3},
	{ChallengeID: 47, MonsterID: 43, Count: 1}, {ChallengeID: 47, MonsterID: 33, Count: 3},
	{ChallengeID: 48, MonsterID: 37, Count: 2}, {ChallengeID: 48, MonsterID: 39, Count: 2},
	{ChallengeID: 49, MonsterID: 44, Count: 1}, {ChallengeID: 49, MonsterID: 38, Count: 2},
	{ChallengeID: 50, MonsterID: 45, Count: 1}, {ChallengeID: 50, MonsterID: 44, Count: 1},
	{ChallengeID: 51, MonsterID: 45, Count: 1}, {ChallengeID: 51, MonsterID: 42, Count: 1}, {ChallengeID: 51, MonsterID: 39, Count: 1},
}

// Goods 商品
var Goods = []model.Goods{
	{ID: 1, Name: "下品淬体丹", Price: 50, Description: "增加五点体魄上限，存在灵力反噬风险,长期服用存在耐药性"},
	{ID: 2, Name: "下品莽牛血", Price: 50, Description: "增加一点攻击，存在灵力反噬风险,长期服用存在耐药性"},
	{ID: 3, Name: "下品玄龟甲", Price: 50, Description: "增加一点防御，存在灵力反噬风险,长期服用存在耐药性"},
	{ID: 4, Name: "下品灵蛇皮", Price: 50, Description: "增加一点速度，存在灵力反噬风险,长期服用存在耐药性"},

	{ID: 5, Name: "逍遥散", Price: 20, Description: "逍遥一念间，天地皆可得，有几率触发顿悟的丹药，可能会得到大量经验"},
	{ID: 6, Name: "修为丹", Price: 20, Description: "增加二十点经验"},
	{ID: 7, Name: "愈伤丹", Price: 20, Description: "瞬间恢复三十点生命值"},
	{ID: 8, Name: "金币罐子", Price: 100, Description: "会获得随机数量的金币 -> Random(1, Max(Level, 150) )"},

	{ID: 9, Name: "上品淬体丹", Price: 5000, Description: "增加十点体魄上限，药效温和非常稳定,甚至可以突破境界的限制"},
	{ID: 10, Name: "上品莽牛血", Price: 5000, Description: "增加两点攻击，药效温和非常稳定,甚至可以突破境界的限制"},
	{ID: 11, Name: "上品玄龟甲", Price: 5000, Description: "增加两点防御，药效温和非常稳定,甚至可以突破境界的限制"},
	{ID: 12, Name: "上品灵蛇皮", Price: 5000, Description: "增加两点速度，药效温和非常稳定,甚至可以突破境界的限制"},
}

var MonsterMap map[uint]model.Monster
var ChallengeMap map[uint]model.Challenge
var ChallengeMonsterMap map[uint]model.ChallengeMonster
var GoodsMap map[uint]model.Goods

func init() {
	MonsterMap = make(map[uint]model.Monster)
	ChallengeMap = make(map[uint]model.Challenge)
	ChallengeMonsterMap = make(map[uint]model.ChallengeMonster)
	GoodsMap = make(map[uint]model.Goods)

	for _, v := range monsters {
		MonsterMap[v.ID] = v
	}
	for _, v := range Challenges {
		ChallengeMap[v.ID] = v
	}
	for ind, v := range challengeMonsters {
		v.Monster = MonsterMap[v.MonsterID]
		ChallengeMonsterMap[uint(ind+1)] = v
	}
	for _, v := range Goods {
		GoodsMap[v.ID] = v
	}

}
