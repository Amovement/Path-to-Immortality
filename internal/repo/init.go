package repo

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"sort"
)

var monsters = []model.Monster{
	// （1-30级）
	{ID: 1, Name: "赤毛灵狐", Hp: 20, HpLimit: 20, Attack: 2, Defense: 2, Speed: 2, Level: 1},
	{ID: 2, Name: "铁皮妖鼠", Hp: 50, HpLimit: 50, Attack: 4, Defense: 4, Speed: 4, Level: 3},
	{ID: 3, Name: "腐气僵尸", Hp: 80, HpLimit: 80, Attack: 6, Defense: 6, Speed: 6, Level: 5},
	{ID: 4, Name: "青面山魈", Hp: 130, HpLimit: 130, Attack: 10, Defense: 8, Speed: 8, Level: 8},
	{ID: 5, Name: "启灵魔仆", Hp: 180, HpLimit: 180, Attack: 14, Defense: 12, Speed: 10, Level: 12},
	{ID: 6, Name: "百年石灵", Hp: 220, HpLimit: 220, Attack: 14, Defense: 18, Speed: 10, Level: 15},
	{ID: 7, Name: "风系小妖", Hp: 220, HpLimit: 220, Attack: 16, Defense: 18, Speed: 20, Level: 18},
	{ID: 8, Name: "无名散修", Hp: 250, HpLimit: 250, Attack: 24, Defense: 18, Speed: 20, Level: 25},

	// （31-60 级）
	{ID: 9, Name: "筑基境傀儡", Hp: 455, HpLimit: 455, Attack: 47, Defense: 39, Speed: 23, Level: 32},
	{ID: 10, Name: "火鸦妖王", Hp: 494, HpLimit: 494, Attack: 57, Defense: 39, Speed: 25, Level: 35},
	{ID: 11, Name: "玄冰精怪", Hp: 572, HpLimit: 572, Attack: 57, Defense: 47, Speed: 25, Level: 40},
	{ID: 12, Name: "金丹长老分身", Hp: 650, HpLimit: 650, Attack: 57, Defense: 52, Speed: 27, Level: 45},
	{ID: 13, Name: "血炼魔徒", Hp: 689, HpLimit: 689, Attack: 65, Defense: 52, Speed: 27, Level: 50},
	{ID: 14, Name: "雷纹巨猿", Hp: 754, HpLimit: 754, Attack: 78, Defense: 52, Speed: 27, Level: 55},

	// （61-90 级）
	{ID: 15, Name: "护法", Hp: 832, HpLimit: 832, Attack: 83, Defense: 57, Speed: 27, Level: 65},
	{ID: 16, Name: "骨甲尸王", Hp: 900, HpLimit: 900, Attack: 83, Defense: 75, Speed: 27, Level: 70},
	{ID: 17, Name: "风灵圣女", Hp: 900, HpLimit: 900, Attack: 91, Defense: 75, Speed: 41, Level: 75},
	{ID: 18, Name: "宗门长老", Hp: 900, HpLimit: 900, Attack: 100, Defense: 75, Speed: 41, Level: 80},
	{ID: 19, Name: "熔岩古龙", Hp: 936, HpLimit: 936, Attack: 130, Defense: 78, Speed: 41, Level: 85},

	// （91-120 级）
	{ID: 20, Name: "上古傀儡", Hp: 1200, HpLimit: 1200, Attack: 150, Defense: 105, Speed: 41, Level: 95},
	{ID: 21, Name: "玄天蛇姬", Hp: 1260, HpLimit: 1260, Attack: 150, Defense: 108, Speed: 43, Level: 100},
	{ID: 22, Name: "血河老怪", Hp: 1335, HpLimit: 1335, Attack: 165, Defense: 111, Speed: 46, Level: 105},
	{ID: 23, Name: "雷劫剑灵", Hp: 1335, HpLimit: 1335, Attack: 186, Defense: 114, Speed: 72, Level: 110},
	{ID: 24, Name: "元婴境大长老", Hp: 1395, HpLimit: 1395, Attack: 189, Defense: 117, Speed: 96, Level: 115},

	// （121-150 级）难度巨大提升
	{ID: 25, Name: "化神妖兽", Hp: 3000, HpLimit: 3000, Attack: 240, Defense: 120, Speed: 96, Level: 125},
	{ID: 26, Name: "空间裂缝兽", Hp: 3000, HpLimit: 3000, Attack: 264, Defense: 123, Speed: 125, Level: 130},
	{ID: 27, Name: "上古战魂", Hp: 3120, HpLimit: 3120, Attack: 300, Defense: 126, Speed: 127, Level: 135},
	{ID: 28, Name: "炼虚长老分身", Hp: 3120, HpLimit: 3120, Attack: 330, Defense: 129, Speed: 130, Level: 140},
	{ID: 29, Name: "混沌石灵", Hp: 3180, HpLimit: 3180, Attack: 360, Defense: 165, Speed: 132, Level: 145},

	// （151-180 级）难度巨大提升
	{ID: 30, Name: "魔将", Hp: 3904, HpLimit: 3904, Attack: 525, Defense: 186, Speed: 146, Level: 155},
	{ID: 31, Name: "九天玄女", Hp: 3904, HpLimit: 3904, Attack: 544, Defense: 192, Speed: 159, Level: 160},
	{ID: 32, Name: "万尸之主", Hp: 4100, HpLimit: 4100, Attack: 547, Defense: 248, Speed: 161, Level: 165},
	{ID: 33, Name: "宗主化身", Hp: 4100, HpLimit: 4100, Attack: 563, Defense: 282, Speed: 164, Level: 170},
	{ID: 34, Name: "洪荒巨兽", Hp: 4290, HpLimit: 4290, Attack: 576, Defense: 307, Speed: 169, Level: 175},

	// （181-210 级）难度巨大提升
	{ID: 35, Name: "合体仙将", Hp: 4685, HpLimit: 4685, Attack: 647, Defense: 403, Speed: 173, Level: 185},
	{ID: 36, Name: "灭世魔魂", Hp: 4685, HpLimit: 4685, Attack: 714, Defense: 374, Speed: 175, Level: 190},
	{ID: 37, Name: "天道守护者", Hp: 4875, HpLimit: 4875, Attack: 721, Defense: 442, Speed: 178, Level: 195},
	{ID: 38, Name: "圣祖", Hp: 5372, HpLimit: 5372, Attack: 734, Defense: 452, Speed: 235, Level: 200},
	{ID: 39, Name: "混沌之影", Hp: 5372, HpLimit: 5372, Attack: 762, Defense: 456, Speed: 238, Level: 205},

	// （211-250 级）难度巨大提升
	{ID: 40, Name: "雷劫之灵", Hp: 5372, HpLimit: 5372, Attack: 789, Defense: 462, Speed: 258, Level: 215},
	{ID: 41, Name: "空间法则兽", Hp: 5372, HpLimit: 5372, Attack: 850, Defense: 468, Speed: 340, Level: 220},
	{ID: 42, Name: "上古仙尊残魂", Hp: 5474, HpLimit: 5474, Attack: 898, Defense: 476, Speed: 347, Level: 225},
	{ID: 43, Name: "渡劫天魔", Hp: 5712, HpLimit: 5712, Attack: 952, Defense: 510, Speed: 354, Level: 230},
	{ID: 44, Name: "天道执法者", Hp: 6000, HpLimit: 6000, Attack: 1034, Defense: 612, Speed: 366, Level: 240},
	{ID: 45, Name: "鸿蒙本源兽", Hp: 6256, HpLimit: 6256, Attack: 1122, Defense: 646, Speed: 374, Level: 250},
}

// Challenges 50个挑战任务（每5级一个梯度，覆盖1-250级）
var Challenges = []model.Challenge{
	{ID: 0, LevelLimit: 30, Title: "每日功课 - 心魔", Gold: 5000}, // 福利局

	{ID: 1, LevelLimit: 30, Title: "新手林地试炼", Gold: 100},   // 1级
	{ID: 2, LevelLimit: 30, Title: "乱葬岗清剿", Gold: 300},    // 5级
	{ID: 3, LevelLimit: 30, Title: "黑风洞探秘", Gold: 500},    // 10级
	{ID: 4, LevelLimit: 30, Title: "灵草园守护", Gold: 800},    // 15级
	{ID: 5, LevelLimit: 30, Title: "引气期修士考验", Gold: 1200}, // 20级
	{ID: 6, LevelLimit: 30, Title: "妖兽潮拦截", Gold: 2000},   // 25级
	{ID: 7, LevelLimit: 30, Title: "引气境圆满挑战", Gold: 3000}, // 30级

	{ID: 8, LevelLimit: 60, Title: "筑基秘境初探", Gold: 5000},    // 35级
	{ID: 9, LevelLimit: 60, Title: "火鸦巢清理", Gold: 6000},     // 40级
	{ID: 10, LevelLimit: 60, Title: "玄冰窟探险", Gold: 7000},    // 45级
	{ID: 11, LevelLimit: 60, Title: "筑基期入门考核", Gold: 8000},  // 50级
	{ID: 12, LevelLimit: 60, Title: "血魔谷围剿", Gold: 9000},    // 55级
	{ID: 13, LevelLimit: 60, Title: "筑基境圆满挑战", Gold: 10000}, // 60级

	{ID: 14, LevelLimit: 90, Title: "金丹洞府开启", Gold: 12000},  // 65级
	{ID: 15, LevelLimit: 90, Title: "尸王殿破局", Gold: 14000},   // 70级
	{ID: 16, LevelLimit: 90, Title: "风灵谷试炼", Gold: 16000},   // 75级
	{ID: 17, LevelLimit: 90, Title: "金丹长老考验", Gold: 18000},  // 80级
	{ID: 18, LevelLimit: 90, Title: "古龙巢穴探秘", Gold: 20000},  // 85级
	{ID: 19, LevelLimit: 90, Title: "金丹境圆满挑战", Gold: 22000}, // 90级

	{ID: 20, LevelLimit: 120, Title: "元婴遗迹探索", Gold: 25000},  // 95级
	{ID: 21, LevelLimit: 120, Title: "玄天蛇窟历险", Gold: 28000},  // 100级
	{ID: 22, LevelLimit: 120, Title: "血河源头镇压", Gold: 31000},  // 105级
	{ID: 23, LevelLimit: 120, Title: "雷劫剑灵降服", Gold: 34000},  // 110级
	{ID: 24, LevelLimit: 120, Title: "元婴大典挑战", Gold: 37000},  // 115级
	{ID: 25, LevelLimit: 120, Title: "元婴境圆满挑战", Gold: 40000}, // 120级

	{ID: 26, LevelLimit: 150, Title: "化神秘境闯关", Gold: 45000},  // 125级
	{ID: 27, LevelLimit: 150, Title: "空间裂缝平定", Gold: 50000},  // 130级
	{ID: 28, LevelLimit: 150, Title: "上古战场历练", Gold: 55000},  // 135级
	{ID: 29, LevelLimit: 150, Title: "化神长老对决", Gold: 60000},  // 140级
	{ID: 30, LevelLimit: 150, Title: "混沌石林试炼", Gold: 65000},  // 145级
	{ID: 31, LevelLimit: 150, Title: "化神境圆满挑战", Gold: 70000}, // 150级

	{ID: 32, LevelLimit: 180, Title: "魔窟征伐", Gold: 77000},     // 155级
	{ID: 33, LevelLimit: 180, Title: "九天瑶池试炼", Gold: 80000},   // 160级
	{ID: 34, LevelLimit: 180, Title: "万尸窟净化", Gold: 87000},    // 165级
	{ID: 35, LevelLimit: 180, Title: "宗主宝座争夺", Gold: 90000},   // 170级
	{ID: 36, LevelLimit: 180, Title: "洪荒兽域探险", Gold: 97000},   // 175级
	{ID: 37, LevelLimit: 180, Title: "炼虚境圆满挑战", Gold: 100000}, // 180级

	{ID: 38, LevelLimit: 210, Title: "仙域开启", Gold: 110000},    // 185级
	{ID: 39, LevelLimit: 210, Title: "灭世魔魂封印", Gold: 120000},  // 190级
	{ID: 40, LevelLimit: 210, Title: "天道试炼场", Gold: 130000},   // 195级
	{ID: 41, LevelLimit: 210, Title: "圣祖传承考验", Gold: 140000},  // 200级
	{ID: 42, LevelLimit: 210, Title: "混沌空间历练", Gold: 150000},  // 205级
	{ID: 43, LevelLimit: 210, Title: "合体境圆满挑战", Gold: 160000}, // 210级

	{ID: 44, LevelLimit: 240, Title: "一重雷劫挑战", Gold: 180000},  // 215级
	{ID: 45, LevelLimit: 240, Title: "空间法则试炼", Gold: 200000},  // 220级
	{ID: 46, LevelLimit: 240, Title: "仙尊传承争夺", Gold: 300000},  // 225级
	{ID: 47, LevelLimit: 240, Title: "天魔巢穴围剿", Gold: 400000},  // 230级
	{ID: 48, LevelLimit: 240, Title: "天道法则领悟", Gold: 500000},  // 235级
	{ID: 49, LevelLimit: 240, Title: "执法者考验", Gold: 600000},   // 240级
	{ID: 50, LevelLimit: 240, Title: "鸿蒙本源试炼", Gold: 700000},  // 245级
	{ID: 51, LevelLimit: 240, Title: "渡劫境圆满挑战", Gold: 800000}, // 250级
}

// 挑战-怪物关联数据（按每级3点属性成长平衡难度）
var challengeMonsters = []model.ChallengeMonster{
	// （1-30级）
	{ChallengeID: 1, MonsterID: 1, Count: 3}, {ChallengeID: 1, MonsterID: 2, Count: 1},
	{ChallengeID: 2, MonsterID: 3, Count: 2}, {ChallengeID: 2, MonsterID: 4, Count: 1},
	{ChallengeID: 3, MonsterID: 4, Count: 3}, {ChallengeID: 3, MonsterID: 5, Count: 1},
	{ChallengeID: 4, MonsterID: 5, Count: 2}, {ChallengeID: 4, MonsterID: 7, Count: 2},
	{ChallengeID: 5, MonsterID: 7, Count: 3}, {ChallengeID: 5, MonsterID: 8, Count: 1},
	{ChallengeID: 6, MonsterID: 4, Count: 5}, {ChallengeID: 6, MonsterID: 5, Count: 3}, {ChallengeID: 6, MonsterID: 8, Count: 2},
	{ChallengeID: 7, MonsterID: 8, Count: 5}, {ChallengeID: 7, MonsterID: 6, Count: 3},

	// （31-60级）
	{ChallengeID: 8, MonsterID: 9, Count: 2}, {ChallengeID: 8, MonsterID: 10, Count: 1},
	{ChallengeID: 9, MonsterID: 10, Count: 3}, {ChallengeID: 9, MonsterID: 11, Count: 1},
	{ChallengeID: 10, MonsterID: 11, Count: 2}, {ChallengeID: 10, MonsterID: 12, Count: 1},
	{ChallengeID: 11, MonsterID: 12, Count: 2}, {ChallengeID: 11, MonsterID: 13, Count: 1},
	{ChallengeID: 12, MonsterID: 13, Count: 3}, {ChallengeID: 12, MonsterID: 14, Count: 1},
	{ChallengeID: 13, MonsterID: 14, Count: 2}, {ChallengeID: 13, MonsterID: 12, Count: 2},

	// （61-90级）
	{ChallengeID: 14, MonsterID: 15, Count: 2}, {ChallengeID: 14, MonsterID: 16, Count: 1},
	{ChallengeID: 15, MonsterID: 16, Count: 2}, {ChallengeID: 15, MonsterID: 17, Count: 1},
	{ChallengeID: 16, MonsterID: 17, Count: 2}, {ChallengeID: 16, MonsterID: 18, Count: 1},
	{ChallengeID: 17, MonsterID: 18, Count: 2}, {ChallengeID: 17, MonsterID: 19, Count: 1},
	{ChallengeID: 18, MonsterID: 19, Count: 2}, {ChallengeID: 18, MonsterID: 18, Count: 1},
	{ChallengeID: 19, MonsterID: 19, Count: 3}, {ChallengeID: 19, MonsterID: 17, Count: 2},

	// （91-120级）
	{ChallengeID: 20, MonsterID: 20, Count: 2}, {ChallengeID: 20, MonsterID: 21, Count: 1},
	{ChallengeID: 21, MonsterID: 21, Count: 2}, {ChallengeID: 21, MonsterID: 22, Count: 1},
	{ChallengeID: 22, MonsterID: 22, Count: 2}, {ChallengeID: 22, MonsterID: 23, Count: 1},
	{ChallengeID: 23, MonsterID: 23, Count: 2}, {ChallengeID: 23, MonsterID: 24, Count: 1},
	{ChallengeID: 24, MonsterID: 24, Count: 2}, {ChallengeID: 24, MonsterID: 22, Count: 1},
	{ChallengeID: 25, MonsterID: 24, Count: 3}, {ChallengeID: 25, MonsterID: 23, Count: 2},

	// （121-150级）
	{ChallengeID: 26, MonsterID: 25, Count: 2}, {ChallengeID: 26, MonsterID: 26, Count: 1},
	{ChallengeID: 27, MonsterID: 26, Count: 2}, {ChallengeID: 27, MonsterID: 27, Count: 1},
	{ChallengeID: 28, MonsterID: 27, Count: 2}, {ChallengeID: 28, MonsterID: 28, Count: 1},
	{ChallengeID: 29, MonsterID: 28, Count: 2}, {ChallengeID: 29, MonsterID: 29, Count: 1},
	{ChallengeID: 30, MonsterID: 29, Count: 2}, {ChallengeID: 30, MonsterID: 27, Count: 1},
	{ChallengeID: 31, MonsterID: 29, Count: 3}, {ChallengeID: 31, MonsterID: 28, Count: 2},

	// （151-180级）
	{ChallengeID: 32, MonsterID: 30, Count: 2}, {ChallengeID: 32, MonsterID: 31, Count: 1},
	{ChallengeID: 33, MonsterID: 31, Count: 2}, {ChallengeID: 33, MonsterID: 32, Count: 1},
	{ChallengeID: 34, MonsterID: 32, Count: 2}, {ChallengeID: 34, MonsterID: 33, Count: 1},
	{ChallengeID: 35, MonsterID: 33, Count: 2}, {ChallengeID: 35, MonsterID: 34, Count: 1},
	{ChallengeID: 36, MonsterID: 34, Count: 2}, {ChallengeID: 36, MonsterID: 32, Count: 1},
	{ChallengeID: 37, MonsterID: 34, Count: 3}, {ChallengeID: 37, MonsterID: 33, Count: 2},

	// （181-210级）
	{ChallengeID: 38, MonsterID: 35, Count: 2}, {ChallengeID: 38, MonsterID: 36, Count: 1},
	{ChallengeID: 39, MonsterID: 36, Count: 2}, {ChallengeID: 39, MonsterID: 37, Count: 1},
	{ChallengeID: 40, MonsterID: 37, Count: 2}, {ChallengeID: 40, MonsterID: 38, Count: 1},
	{ChallengeID: 41, MonsterID: 38, Count: 2}, {ChallengeID: 41, MonsterID: 39, Count: 1},
	{ChallengeID: 42, MonsterID: 39, Count: 2}, {ChallengeID: 42, MonsterID: 37, Count: 1},
	{ChallengeID: 43, MonsterID: 39, Count: 3}, {ChallengeID: 43, MonsterID: 38, Count: 2},

	// （211-250级）
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
	{UUid: 1, Name: "下品淬体丹", Type: model.ItemTypeConsume, Price: 50, Description: "增加五点体魄上限，存在灵力反噬风险,长期服用存在耐药性"},
	{UUid: 2, Name: "下品莽牛血", Type: model.ItemTypeConsume, Price: 50, Description: "增加一点攻击，存在灵力反噬风险,长期服用存在耐药性"},
	{UUid: 3, Name: "下品玄龟甲", Type: model.ItemTypeConsume, Price: 50, Description: "增加一点防御，存在灵力反噬风险,长期服用存在耐药性"},
	{UUid: 4, Name: "下品灵蛇皮", Type: model.ItemTypeConsume, Price: 50, Description: "增加一点速度，存在灵力反噬风险,长期服用存在耐药性"},
	{UUid: 5, Name: "逍遥散", Type: model.ItemTypeConsume, Price: 20, Description: "逍遥一念间，天地皆可得，有几率触发顿悟的丹药，可能会得到大量经验"},
	{UUid: 6, Name: "下品修为丹", Type: model.ItemTypeConsume, Price: 50, Description: "对金丹境及以下修士效果比较好的丹药, 可增加 10 点修炼经验, 高境界效果骤减"},
	{UUid: 7, Name: "愈伤丹", Type: model.ItemTypeConsume, Price: 20, Description: "瞬间恢复十五点体魄值"},
	{UUid: 8, Name: "金币罐子", Type: model.ItemTypeConsume, Price: 100, Description: "会获得随机数量的金币 -> Random(1, Max(Level, 150) )"},
	{UUid: 9, Name: "上品淬体丹", Type: model.ItemTypeConsume, Price: 5000, Description: "增加十点体魄上限，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 10, Name: "上品莽牛血", Type: model.ItemTypeConsume, Price: 5000, Description: "增加两点攻击，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 11, Name: "上品玄龟甲", Type: model.ItemTypeConsume, Price: 5000, Description: "增加两点防御，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 12, Name: "上品灵蛇皮", Type: model.ItemTypeConsume, Price: 5000, Description: "增加两点速度，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 13, Name: "混沌清浊气", Type: model.ItemTypeConsume, Price: 50000, Description: "会让体内的潜能躁动起来，获得一点新的潜能点，对轮回转世之人有更好的效果"},
	{UUid: 14, Name: "玄晶", Type: model.ItemTypeMaterial, Price: 2500, Description: "亮晶晶的矿物, 配合上`精魄`即可打造一柄随机的法器"},
	{UUid: 15, Name: "精魄", Type: model.ItemTypeMaterial, Price: 5000, Description: "它似乎还活着, 配合上`玄晶`即可打造一柄随机的法器"},
	{UUid: 16, Name: "锻铁", Type: model.ItemTypeMaterial, Price: 1000, Description: "内部充满能量的铁块, 可以用来提升法器等级"},
	{UUid: 17, Name: "上品修为丹", Type: model.ItemTypeConsume, Price: 400, Description: "对炼虚境及以下修士效果比较好的丹药, 可增加 10 点修炼经验, 高境界效果骤减"},
}

const (
	XiaPinCuiTiDanUUid     = iota + 1 // 下品淬体丹
	XiaPinMangNiuXueUUid              // 下品莽牛血
	XiaPinXuanGuiJiaUUid              // 下品玄龟甲
	XiaPinLingShePiUUid               // 下品灵蛇皮
	XiaoYaoSanUUid                    // 逍遥散
	XiaPinXiuWeiDanUUid               // 修为丹
	YuShangDanUUid                    // 愈伤丹
	JinBiGuanZiUUid                   // 金币罐子
	ShangPinCuiTiDanUUid              // 上品淬体丹
	ShangPinMangNiuXueUUid            // 上品莽牛血
	ShangPinXuanGuiJiaUUid            // 上品玄龟甲
	ShangPinLingShePiUUid             // 上品灵蛇皮
	HunDunQingZhuoQiUUid              // 混沌清浊气
	XuanJingUUid                      // 玄晶
	JingPoUUid                        // 精魄
	DuanTieUUid                       // 锻铁
	ShangPinXiuWeiDanUUid             // 上品修为丹
)

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
		GoodsMap[v.UUid] = v
	}
	// sort goods by price
	sort.Slice(Goods, func(i, j int) bool {
		if Goods[i].Price == Goods[j].Price {
			return Goods[i].UUid > Goods[j].UUid
		}
		return Goods[i].Price < Goods[j].Price
	})

}
