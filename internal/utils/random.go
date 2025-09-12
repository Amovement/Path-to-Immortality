package utils

import (
	"math/rand"
	"time"
)

func GetRandomInt64(min, max int64) int64 {
	return min + int64(rand.Intn(int(max-min)))
}

func GetRandomMinutes(min, max int) time.Duration {
	if max <= min {
		max = min + 1
	}
	return time.Duration(min+rand.Intn(max-min)) * time.Minute
}
