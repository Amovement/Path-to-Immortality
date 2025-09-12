package model

type Monster struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Hp      int64  `json:"hp"`
	HpLimit int64  `json:"hpLimit"`
	Attack  int64  `json:"attack"`
	Defense int64  `json:"defense"`
	Speed   int64  `json:"speed"`
	Level   int64  `json:"level"`
}
