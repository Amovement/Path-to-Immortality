package test

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"testing"
)

func TestRandomEquip(t *testing.T) {
	t.Logf("%+v", model.RandomEquip(0, 10, -1))
	t.Logf("%+v", model.RandomEquip(10, 20, -1))
	t.Logf("%+v", model.RandomEquip(20, 30, -1))
	t.Logf("%+v", model.RandomEquip(30, 40, -1))
	t.Logf("%+v", model.RandomEquip(40, 50, -1))
}
