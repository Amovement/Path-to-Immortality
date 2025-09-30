package model

type Goods struct {
	UUid        uint   `json:"uuid"` // 全局唯一编号
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Type        uint   `json:"type"` // 0:消耗品 1:装备 2:材料
}
