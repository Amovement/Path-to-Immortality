package model

type Goods struct {
	UUid        uint   `json:"uuid"` // 全局唯一编号
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
