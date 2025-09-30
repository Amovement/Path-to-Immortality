package register

import (
	"github.com/Amovement/Path-to-Immortality-WASM/internal"
	"syscall/js"
)

func RegisterGoodsCallbacks(core *internal.Core) {
	js.Global().Set("getGoodsList", js.FuncOf(core.GetGoodsList))
	js.Global().Set("buyGoods", js.FuncOf(core.BuyGoods))
}

func RegisterUserCallbacks(core *internal.Core) {
	js.Global().Set("getUserInfo", js.FuncOf(core.GetUserInfo))
	js.Global().Set("setUsername", js.FuncOf(core.SetUsername))
	js.Global().Set("allocate", js.FuncOf(core.Allocate))
	js.Global().Set("heal", js.FuncOf(core.Heal))
	js.Global().Set("cultivation", js.FuncOf(core.Cultivation))
	js.Global().Set("getGold", js.FuncOf(core.GetGold))
	js.Global().Set("restart", js.FuncOf(core.Restart))
}

func RegisterChallengeCallbacks(core *internal.Core) {
	js.Global().Set("listChallenge", js.FuncOf(core.ListChallenge))
	js.Global().Set("joinChallenge", js.FuncOf(core.JoinChallenge))
}

func RegisterVersionCallbacks(core *internal.Core) {
	js.Global().Set("getVersion", js.FuncOf(core.GetVersion))
}

func RegisterBagCallbacks(core *internal.Core) {
	js.Global().Set("getBag", js.FuncOf(core.GetBag))
	js.Global().Set("useBagItem", js.FuncOf(core.UseBagItem))
}

func RegisterEquipCallbacks(core *internal.Core) {
	js.Global().Set("takeOffEquip", js.FuncOf(core.TakeOffEquip))
	js.Global().Set("getUserEquipAttributes", js.FuncOf(core.GetUserEquipAttributes))
}
