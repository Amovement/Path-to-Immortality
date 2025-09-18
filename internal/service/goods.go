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
	return "购买成功"
}

func (s *GoodsService) GetGoodsList() []model.Goods {
	return s.Goods
}
