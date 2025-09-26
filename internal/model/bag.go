package model

type Bag struct {
	Items []*Item `json:"Items"`
}

func NewBag() *Bag {
	return &Bag{
		Items: []*Item{},
	}
}

const (
	BagStorageKey = "_Path_2_Immortality_Bag_"
)
