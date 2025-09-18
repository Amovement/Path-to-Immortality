//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal"
	"syscall/js"
)

func registerUserCallbacks(core *internal.Core) {
	js.Global().Set("getUserInfo", js.FuncOf(core.GetUserInfo))
	js.Global().Set("setUsername", js.FuncOf(core.SetUsername))
	js.Global().Set("allocate", js.FuncOf(core.Allocate))
	js.Global().Set("heal", js.FuncOf(core.Heal))
	js.Global().Set("cultivation", js.FuncOf(core.Cultivation))
}

func registerChallengeCallbacks(core *internal.Core) {
	js.Global().Set("listChallenge", js.FuncOf(core.ListChallenge))
	js.Global().Set("joinChallenge", js.FuncOf(core.JoinChallenge))
}

func registerVersionCallbacks(core *internal.Core) {
	js.Global().Set("getVersion", js.FuncOf(core.GetVersion))
}

func main() {
	core := internal.NewCore()
	registerUserCallbacks(core)
	registerChallengeCallbacks(core)
	registerVersionCallbacks(core)
	<-make(chan struct{}) // 保持运行
}
