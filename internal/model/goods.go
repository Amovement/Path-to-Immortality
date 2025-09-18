package model

type Goods struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
