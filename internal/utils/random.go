package utils

import (
	"math/rand"
	"time"
)

// GetRandomInt64 generates a random int64 number within the specified range [min, max).
// Parameters:
//
//	min - the minimum value (inclusive)
//	max - the maximum value (exclusive)
//
// Returns:
//
//	a random int64 number in the range [min, max)
func GetRandomInt64(min, max int64) int64 {
	return min + int64(rand.Intn(int(max-min)))
}

func GetRandomMinutes(min, max int) time.Duration {
	if max <= min {
		max = min + 1
	}
	return time.Duration(min+rand.Intn(max-min)) * time.Minute
}
