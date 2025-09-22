package service

import (
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
)

type GoodsService struct {
	GoodsMap map[uint]model.Goods
	Goods    []model.Goods
}

func NewGoodsService() *GoodsService {
	return &GoodsService{
		GoodsMap: repo.GetGoodsMap(),
		Goods:    repo.GetGoodsList(),
	}
}

func (s *GoodsService) BuyGoods(goodsId int) string {
	goods, ok := s.GoodsMap[uint(goodsId)]
	if !ok {
		return "商品不存在"
	}
	user := getLocalUser()
	if user.Gold < goods.Price {
		return "金币不足"
	}
	user.Gold -= goods.Price
	msg := "购买成功! "
	//	{ID: 1, Name: "下品淬体丹", Price: 50, Description: "增加五点体魄上限，存在灵力反噬风险,长期服用存在耐药性"},
	//	{ID: 2, Name: "下品莽牛血", Price: 50, Description: "增加一点攻击，存在灵力反噬风险,长期服用存在耐药性"},
	//	{ID: 3, Name: "下品玄龟甲", Price: 50, Description: "增加一点防御，存在灵力反噬风险,长期服用存在耐药性"},
	//	{ID: 4, Name: "下品灵蛇皮", Price: 50, Description: "增加一点速度，存在灵力反噬风险,长期服用存在耐药性"},
	//
	//	{ID: 5, Name: "逍遥散", Price: 20, Description: "逍遥一念间，天地皆可得，有几率触发顿悟的丹药，可能会得到大量经验"},
	//	{ID: 6, Name: "修为丹", Price: 20, Description: "增加二十点经验"},
	//	{ID: 7, Name: "愈伤丹", Price: 20, Description: "瞬间恢复三十点生命值"},
	//	{ID: 8, Name: "金币罐子", Price: 100, Description: "会获得随机数量的金币 -> Random(1, Max(Level, 150) )"},
	//
	//	{ID: 9, Name: "上品淬体丹", Price: 5000, Description: "增加十点体魄上限，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{ID: 10, Name: "上品莽牛血", Price: 5000, Description: "增加两点攻击，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{ID: 11, Name: "上品玄龟甲", Price: 5000, Description: "增加两点防御，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{ID: 12, Name: "上品灵蛇皮", Price: 5000, Description: "增加两点速度，药效温和非常稳定,可以长期服用,但仍有限制"},
	//
	//	{ID: 13, Name: "混沌清浊气", Price: 50000, Description: "会让体内的潜能躁动起来，获得一点新的潜能点"},

	// 检查耐药性
	ok = s.checkGoodsLimit(user, goodsId)
	if !ok {
		msg = msg + " 当前境界服用此物已经达到瓶颈.还是提升了境界再来吧..."
		return msg
	}

	// 随机灵力反噬
	randInt := utils.GetRandomInt64(1, user.Level+100)
	badThingHappened := false

	switch goodsId {
	case 1:
		msg = msg + "你的体魄上限小幅度增强了."
		user.HpLimit += 5
		if randInt <= 70 {
			badThingHappened = true
		}
	case 2:
		msg = msg + "你的攻击力小幅度提升了."
		user.Attack += 1
		if randInt <= 70 {
			badThingHappened = true
		}
	case 3:
		msg = msg + "你的防御力小幅度提升了."
		user.Defense += 1
		if randInt <= 70 {
			badThingHappened = true
		}
	case 4:
		msg = msg + "你的速度小幅度提升了."
		user.Speed += 1
		if randInt <= 70 {
			badThingHappened = true
		}
	case 5:
		msg = msg + " 你陷入了昏迷."
		if randInt == 50 {
			user.Exp = user.Level*10 - 5
			msg = msg + " 你好像明白了什么，突破只差临门一脚. "
		} else {
			msg = msg + " 什么都没有发生."
		}
	case 6:
		user.Exp += 20
		msg = msg + " 获得经验 20 点"
	case 7:
		user.Hp += 30
		if user.Hp > user.HpLimit {
			user.Hp = user.HpLimit
		}
		msg = msg + " 获得生命值 30 点"
	case 8:
		gold := utils.GetRandomInt64(1, utils.Max(user.Level, 150))
		user.Gold += gold
		msg = msg + "好运连连 获得金币 " + fmt.Sprint(gold) + "枚. "

	case 9:
		msg = msg + "气息更加绵长了."
		user.HpLimit += 10
	case 10:
		msg = msg + "攻击变得更为凌厉了."
		user.Attack += 2
	case 11:
		msg = msg + "叠甲, 过."
		user.Defense += 2
	case 12:
		msg = msg + "脚下生风!"
		user.Speed += 2
	case 13:
		msg = msg + "体内的力量涌出来了. 潜能 +1 ."
		user.Potential += 1
	}

	if badThingHappened { // 惩罚
		msg = msg + " 身体内的灵力暴走了!!!"
		user.Hp -= user.HpLimit / 10
		if user.Hp <= 0 {
			user.Hp = 1
		}
		msg = msg + " 损失 " + fmt.Sprint(user.HpLimit/10) + " 点生命值. "
		randInt = utils.GetRandomInt64(1, 5)
		if randInt == 1 {
			user.HpLimit -= 6
			if user.HpLimit <= 10 {
				user.HpLimit = 10
			}
			msg = msg + " 体魄上限扣除了. "
		} else if randInt == 2 {
			user.Attack -= 2
			if user.Attack < 0 {
				user.Attack = 0
			}
			msg = msg + " 攻击力下降了. "
		} else if randInt == 3 {
			user.Defense -= 2
			if user.Defense < 0 {
				user.Defense = 0
			}
			msg = msg + " 防御力下降了. "
		} else if randInt == 4 {
			user.Speed -= 2
			if user.Speed < 0 {
				user.Speed = 0
			}
			msg = msg + " 速度下降了. "
		}
	}

	updateUserInfo(user)
	return msg
}

func (s *GoodsService) checkGoodsLimit(user *model.User, goodsId int) bool {
	switch goodsId {
	case 1:
		if user.HpLimit+5 > user.Level*3*5+10 {
			return false
		}
	case 2:
		if user.Attack+1 > user.Level*3 {
			return false
		}
	case 3:
		if user.Defense+1 > user.Level*3 {
			return false
		}
	case 4:
		if user.Speed+1 > user.Level*3 {
			return false
		}
	}

	if goodsId >= 9 && goodsId <= 12 {
		// 高级的药 需要检查总属性不超过等级的12倍
		totalStat := user.Attack + user.Defense + user.Speed + user.HpLimit/5 + user.Potential // 玩家的总属性
		if totalStat >= user.Level*12 {
			return false
		}
	}

	return true
}

func (s *GoodsService) GetGoodsList() []model.Goods {
	return s.Goods
}
