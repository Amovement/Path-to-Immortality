package test

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"testing"
)

func TestGetGoodsList(t *testing.T) {
	list := repo.GetGoodsList()
	t.Logf("%+v", list)
}
