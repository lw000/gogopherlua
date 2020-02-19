package export_wsclient

import (
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"open": open,
	"send": send,
}

func checkClient(L *lua.LState) *FastWsClient {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*FastWsClient); ok {
		return v
	}
	L.ArgError(1, "wsclient expected")
	return nil
}
