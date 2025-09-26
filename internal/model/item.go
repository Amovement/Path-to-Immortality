package model

type Item struct {
	UUid        int64  `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Count       int64  `json:"count"`
	Type        uint   `json:"type"` // 类型 0 消耗品 1 装备
}

const (
	ItemTypeConsume = 0 // 消耗品
	ItemTypeEquip   = 1 // 装备
)
