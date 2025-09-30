package model

type Bag struct {
	Items      []*Item `json:"Items"`
	RandomUUid int64   `json:"RandomUUid"` // 随机生成物品的 uuid 下标，每次使用后需要更新
}

func NewBag() *Bag {
	return &Bag{
		Items: []*Item{},
	}
}

const (
	BagStorageKey = "_Path_2_Immortality_Bag_"
)
