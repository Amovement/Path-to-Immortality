package model

type Item struct {
	UUid        int64  `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Count       int64  `json:"count"`
	Type        uint   `json:"type"`      // 类型 0 消耗品 1 装备 2 材料
	EquipInfo   *Equip `json:"equipInfo"` // 为装备时才有值
	Status      uint   `json:"status"`    // 状态 0 无状态 1 装备中 目前只有法器会是1
}

const (
	ItemTypeConsume  = 0 // 消耗品
	ItemTypeEquip    = 1 // 装备
	ItemTypeMaterial = 2 //材料
)
