package repo

import "github.com/Amovement/Path-to-Immortality-WASM/internal/model"

func GetMonsterMap() map[uint]model.Monster {
	return MonsterMap
}
func GetChallengeMap() map[uint]model.Challenge {
	return ChallengeMap
}
func GetChallengeMonsterMap() map[uint]model.ChallengeMonster {
	return ChallengeMonsterMap
}

func GetChallengeList() []model.Challenge {
	return Challenges
}

func GetGoodsList() []model.Goods {
	return Goods
}

func GetGoodsMap() map[uint]model.Goods {
	return GoodsMap
}
