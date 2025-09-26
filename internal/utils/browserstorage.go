package utils

import (
	"fmt"
	"syscall/js"
)

// GetStorage 从localStorage读取数据
func GetStorage(key string) (string, bool) {
	localStorage := js.Global().Get("localStorage")

	// 检查localStorage是否可用
	if !localStorage.Truthy() {
		return "", false
	}

	// 调用localStorage.getItem(key)
	value := localStorage.Call("getItem", key)

	// 检查是否存在该键
	if value.IsUndefined() || value.IsNull() {
		return "", false
	}

	return value.String(), true
}

// SetStorage 向localStorage写入数据
func SetStorage(key, value string) bool {
	localStorage := js.Global().Get("localStorage")

	if !localStorage.Truthy() {
		fmt.Println("localStorage is not available")
		return false
	}

	// 调用localStorage.setItem(key, value)
	localStorage.Call("setItem", key, value)
	return true
}

// DeleteStorage 从localStorage删除数据
func DeleteStorage(key string) {
	localStorage := js.Global().Get("localStorage")
	if localStorage.Truthy() {
		localStorage.Call("removeItem", key)
	}
}
