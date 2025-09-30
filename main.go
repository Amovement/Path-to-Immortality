package main

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/register"
)

func main() {
	core := internal.NewCore()
	register.RegisterUserCallbacks(core)
	register.RegisterChallengeCallbacks(core)
	register.RegisterVersionCallbacks(core)
	register.RegisterGoodsCallbacks(core)
	register.RegisterBagCallbacks(core)
	register.RegisterEquipCallbacks(core)
	<-make(chan struct{}) // 保持运行
}
