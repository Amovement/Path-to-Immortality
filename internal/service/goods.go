package service

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
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
	//	{UUid: 1, Name: "下品淬体丹", Price: 50, Description: "增加五点体魄上限，存在灵力反噬风险,长期服用存在耐药性"},
	//	{UUid: 2, Name: "下品莽牛血", Price: 50, Description: "增加一点攻击，存在灵力反噬风险,长期服用存在耐药性"},
	//	{UUid: 3, Name: "下品玄龟甲", Price: 50, Description: "增加一点防御，存在灵力反噬风险,长期服用存在耐药性"},
	//	{UUid: 4, Name: "下品灵蛇皮", Price: 50, Description: "增加一点速度，存在灵力反噬风险,长期服用存在耐药性"},
	//	{UUid: 5, Name: "逍遥散", Price: 20, Description: "逍遥一念间，天地皆可得，有几率触发顿悟的丹药，可能会得到大量经验"},
	//	{UUid: 6, Name: "修为丹", Price: 20, Description: "增加十点经验"},
	//	{UUid: 7, Name: "愈伤丹", Price: 20, Description: "瞬间恢复十五点生命值"},
	//	{UUid: 8, Name: "金币罐子", Price: 100, Description: "会获得随机数量的金币 -> Random(1, Max(Level, 150) )"},
	//	{UUid: 9, Name: "上品淬体丹", Price: 5000, Description: "增加十点体魄上限，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{UUid: 10, Name: "上品莽牛血", Price: 5000, Description: "增加两点攻击，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{UUid: 11, Name: "上品玄龟甲", Price: 5000, Description: "增加两点防御，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{UUid: 12, Name: "上品灵蛇皮", Price: 5000, Description: "增加两点速度，药效温和非常稳定,可以长期服用,但仍有限制"},
	//	{UUid: 13, Name: "混沌清浊气", Price: 50000, Description: "会让体内的潜能躁动起来，获得一点新的潜能点，对轮回转世之人有更好的效果"},

	addBagItem(&model.Item{
		UUid:        int64(goods.UUid),
		Name:        goods.Name,
		Description: goods.Description,
		Price:       goods.Price,
		Count:       1,
		Type:        goods.Type,
	})
	updateUserInfo(user)

	return msg

}

func (s *GoodsService) GetGoodsList() []model.Goods {
	return s.Goods
}
