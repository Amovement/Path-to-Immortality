package model

// Challenge 游戏内的挑战
type Challenge struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Gold       int64  `json:"gold"`
	LevelLimit int64  `json:"levelLimit"` // 挑战的等级限制
}
