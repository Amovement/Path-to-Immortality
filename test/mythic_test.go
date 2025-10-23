package test

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"testing"
)

func TestRandomMonster(t *testing.T) {
	for i := 0; i < 10; i++ {
		mon := model.GenerateRandomMythicMonster(int64(i))
		t.Logf("%+v", mon)
	}
	mon := model.GenerateRandomMythicMonster(int64(50))
	t.Logf("%+v", mon)
	mon = model.GenerateRandomMythicMonster(int64(100))
	t.Logf("%+v", mon)
	mon = model.GenerateRandomMythicMonster(int64(250))
	t.Logf("%+v", mon)
}
