package export_melody

import (
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"handleConnect":       handleConnect,
	"handleMessageBinary": handleMessageBinary,
	"handleDisconnect":    handleDisconnect,
	"handleError":         handleError,
}

func checkClient(L *lua.LState) *SafeMelody {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*SafeMelody); ok {
		return v
	}
	L.ArgError(1, "melody server expected")
	return nil
}

func handleRequest(L *lua.LState) int {

	return 0
}

func handleConnect(L *lua.LState) int {
	return 0
}

func handleDisconnect(L *lua.LState) int {
	return 0
}

func handleMessageBinary(L *lua.LState) int {
	return 0
}

func handleError(L *lua.LState) int {
	return 0
}
