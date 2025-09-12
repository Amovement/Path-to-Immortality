package model

type ChallengeMonster struct {
	ID          uint  `json:"id"`
	ChallengeID uint  `json:"challengeId"`
	MonsterID   uint  `json:"monsterId"`
	Count       int64 `json:"count"` // 数量

	Monster Monster `json:"monster"`
}
