package test

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal/service"
	"testing"
)

func TestListChallenge(t *testing.T) {
	core := service.NewChallengeService()
	list, err := core.ListChallenge()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}
