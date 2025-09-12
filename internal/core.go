//go:build js && wasm
// +build js,wasm

package internal

import (
	"encoding/json"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/service"
	"syscall/js"
)

type Core struct {
	ChallengeService *service.ChallengeService
	UserService      *service.UserService
}

func NewCore() *Core {
	return &Core{
		ChallengeService: service.NewChallengeService(),
		UserService:      service.NewUserService(),
	}
}

// ------------ 挑战类 ---------------------
func (c *Core) ListChallenge(this js.Value, args []js.Value) interface{} {
	challengeList, err := c.ChallengeService.ListChallenge()
	if err != nil {
		return js.Global().Get("Error").New(err.Error())
	}
	bytesData, _ := json.Marshal(challengeList)

	// 关键修复：必须用js.ValueOf()包装返回值
	return string(bytesData)
}

func (c *Core) JoinChallenge(this js.Value, args []js.Value) interface{} {
	challengeId := args[0].Int()
	msg, log := c.ChallengeService.JoinChallenge(challengeId)
	return map[string]interface{}{
		"msg": msg,
		"log": log,
	}
}

// ------------- user 类 -----------------
func (c *Core) GetUserInfo(this js.Value, args []js.Value) interface{} {
	userInfo := c.UserService.GetUserInfo()
	bytesData, _ := json.Marshal(userInfo)

	// 关键修复：必须用js.ValueOf()包装返回值
	return string(bytesData)
}

func (c *Core) SetUsername(this js.Value, args []js.Value) interface{} {
	username := args[0].String()
	c.UserService.SetUsername(username)
	return nil
}

func (c *Core) Allocate(this js.Value, args []js.Value) interface{} {
	stat := args[0].String()
	resp := c.UserService.Allocate(stat)
	return resp
}

func (c *Core) Heal(this js.Value, args []js.Value) interface{} {
	resp := c.UserService.Heal()
	return resp
}
func (c *Core) Cultivation(this js.Value, args []js.Value) interface{} {
	resp := c.UserService.Cultivation()
	return resp
}
