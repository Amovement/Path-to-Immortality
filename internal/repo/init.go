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
	{ID: 6, Name: "百年石灵", Hp: 220, HpLimit: 220, Attack: 12, Defense: 18, Speed: 4, Level: 15},
	{ID: 7, Name: "风系小妖", Hp: 170, HpLimit: 170, Attack: 16, Defense: 6, Speed: 20, Level: 18},
	{ID: 8, Name: "无名散修", Hp: 250, HpLimit: 250, Attack: 24, Defense: 12, Speed: 14, Level: 25},

	// （31-60 级）
	{ID: 9, Name: "筑基境傀儡", Hp: 350, HpLimit: 350, Attack: 36, Defense: 30, Speed: 16, Level: 32},
	{ID: 10, Name: "火鸦妖王", Hp: 380, HpLimit: 380, Attack: 44, Defense: 20, Speed: 22, Level: 35},
	{ID: 11, Name: "玄冰精怪", Hp: 440, HpLimit: 440, Attack: 32, Defense: 36, Speed: 12, Level: 40},
	{ID: 12, Name: "金丹长老分身", Hp: 480, HpLimit: 480, Attack: 40, Defense: 32, Speed: 20, Level: 45},
	{ID: 13, Name: "血炼魔徒", Hp: 530, HpLimit: 530, Attack: 50, Defense: 28, Speed: 16, Level: 50},
	{ID: 14, Name: "雷纹巨猿", Hp: 580, HpLimit: 580, Attack: 60, Defense: 40, Speed: 12, Level: 55},

	// （61-90 级）
	{ID: 15, Name: "护法", Hp: 640, HpLimit: 640, Attack: 64, Defense: 44, Speed: 24, Level: 65},
	{ID: 16, Name: "骨甲尸王", Hp: 690, HpLimit: 690, Attack: 56, Defense: 58, Speed: 16, Level: 70},
	{ID: 17, Name: "风灵圣女", Hp: 620, HpLimit: 620, Attack: 70, Defense: 30, Speed: 34, Level: 75},
	{ID: 18, Name: "宗门长老", Hp: 670, HpLimit: 670, Attack: 76, Defense: 40, Speed: 26, Level: 80},
	{ID: 19, Name: "熔岩古龙", Hp: 720, HpLimit: 720, Attack: 100, Defense: 60, Speed: 14, Level: 85},

	// （91-120 级）
	{ID: 20, Name: "上古傀儡", Hp: 800, HpLimit: 800, Attack: 90, Defense: 70, Speed: 22, Level: 95},
	{ID: 21, Name: "玄天蛇姬", Hp: 840, HpLimit: 840, Attack: 100, Defense: 44, Speed: 36, Level: 100},
	{ID: 22, Name: "血河老怪", Hp: 890, HpLimit: 890, Attack: 110, Defense: 50, Speed: 26, Level: 105},
	{ID: 23, Name: "雷劫剑灵", Hp: 780, HpLimit: 780, Attack: 124, Defense: 36, Speed: 60, Level: 110},
	{ID: 24, Name: "元婴境大长老", Hp: 930, HpLimit: 930, Attack: 116, Defense: 60, Speed: 90, Level: 115},

	// （121-150 级） 难度巨大提升
	{ID: 25, Name: "化神妖兽", Hp: 2000, HpLimit: 2000, Attack: 160, Defense: 76, Speed: 68, Level: 125},
	{ID: 26, Name: "空间裂缝兽", Hp: 1920, HpLimit: 1920, Attack: 176, Defense: 60, Speed: 104, Level: 130},
	{ID: 27, Name: "上古战魂", Hp: 2080, HpLimit: 2080, Attack: 200, Defense: 80, Speed: 72, Level: 135},
	{ID: 28, Name: "炼虚长老分身", Hp: 2000, HpLimit: 2000, Attack: 220, Defense: 84, Speed: 80, Level: 140},
	{ID: 29, Name: "混沌石灵", Hp: 2120, HpLimit: 2120, Attack: 240, Defense: 110, Speed: 48, Level: 145},

	// （151-180 级） 难度巨大提升
	{ID: 30, Name: "魔将", Hp: 2440, HpLimit: 2440, Attack: 328, Defense: 116, Speed: 50, Level: 155},
	{ID: 31, Name: "九天玄女", Hp: 2360, HpLimit: 2360, Attack: 340, Defense: 120, Speed: 124, Level: 160},
	{ID: 32, Name: "万尸之主", Hp: 2560, HpLimit: 2560, Attack: 320, Defense: 156, Speed: 72, Level: 165},
	{ID: 33, Name: "宗主化身", Hp: 2480, HpLimit: 2480, Attack: 352, Defense: 176, Speed: 96, Level: 170},
	{ID: 34, Name: "洪荒巨兽", Hp: 2680, HpLimit: 2680, Attack: 360, Defense: 192, Speed: 60, Level: 175},

	// （181-210 级） 难度巨大提升
	{ID: 35, Name: "合体仙将", Hp: 2840, HpLimit: 2840, Attack: 392, Defense: 244, Speed: 88, Level: 185},
	{ID: 36, Name: "灭世魔魂", Hp: 2780, HpLimit: 2780, Attack: 420, Defense: 220, Speed: 112, Level: 190},
	{ID: 37, Name: "天道守护者", Hp: 2960, HpLimit: 2960, Attack: 400, Defense: 260, Speed: 84, Level: 195},
	{ID: 38, Name: "圣祖", Hp: 3160, HpLimit: 3160, Attack: 432, Defense: 266, Speed: 108, Level: 200},
	{ID: 39, Name: "混沌之影", Hp: 2920, HpLimit: 2920, Attack: 448, Defense: 268, Speed: 140, Level: 205},

	// （211-250 级） 难度巨大提升
	{ID: 40, Name: "雷劫之灵", Hp: 3120, HpLimit: 3120, Attack: 464, Defense: 272, Speed: 152, Level: 215},
	{ID: 41, Name: "空间法则兽", Hp: 3040, HpLimit: 3040, Attack: 500, Defense: 244, Speed: 200, Level: 220},
	{ID: 42, Name: "上古仙尊残魂", Hp: 3220, HpLimit: 3220, Attack: 528, Defense: 280, Speed: 156, Level: 225},
	{ID: 43, Name: "渡劫天魔", Hp: 3360, HpLimit: 3360, Attack: 560, Defense: 300, Speed: 172, Level: 230},
	{ID: 44, Name: "天道执法者", Hp: 3520, HpLimit: 3520, Attack: 608, Defense: 360, Speed: 192, Level: 240},
	{ID: 45, Name: "鸿蒙本源兽", Hp: 3680, HpLimit: 3680, Attack: 660, Defense: 380, Speed: 220, Level: 250},
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
	{UUid: 6, Name: "修为丹", Type: model.ItemTypeConsume, Price: 20, Description: "增加十点经验"},
	{UUid: 7, Name: "愈伤丹", Type: model.ItemTypeConsume, Price: 20, Description: "瞬间恢复十五点生命值"},
	{UUid: 8, Name: "金币罐子", Type: model.ItemTypeConsume, Price: 100, Description: "会获得随机数量的金币 -> Random(1, Max(Level, 150) )"},
	{UUid: 9, Name: "上品淬体丹", Type: model.ItemTypeConsume, Price: 5000, Description: "增加十点体魄上限，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 10, Name: "上品莽牛血", Type: model.ItemTypeConsume, Price: 5000, Description: "增加两点攻击，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 11, Name: "上品玄龟甲", Type: model.ItemTypeConsume, Price: 5000, Description: "增加两点防御，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 12, Name: "上品灵蛇皮", Type: model.ItemTypeConsume, Price: 5000, Description: "增加两点速度，药效温和非常稳定,可以长期服用,但仍有限制"},
	{UUid: 13, Name: "混沌清浊气", Type: model.ItemTypeConsume, Price: 50000, Description: "会让体内的潜能躁动起来，获得一点新的潜能点，对轮回转世之人有更好的效果"},

	{UUid: 14, Name: "玄晶", Type: model.ItemTypeMaterial, Price: 2500, Description: "亮晶晶的矿物, 配合上`精魄`即可打造一柄随机的法器"},
	{UUid: 15, Name: "精魄", Type: model.ItemTypeMaterial, Price: 5000, Description: "它似乎还活着, 配合上`玄晶`即可打造一柄随机的法器"},
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
