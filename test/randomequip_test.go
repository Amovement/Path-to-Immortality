package test

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"testing"
)

func TestRandomEquip(t *testing.T) {
	t.Logf("%+v", model.RandomEquip())
}
