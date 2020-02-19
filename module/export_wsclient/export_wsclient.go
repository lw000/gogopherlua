package export_wsclient

import (
	lua "github.com/yuin/gopher-lua"
)

const (
	SERVICE_TYPENAME = "wsclient{service}"
)

var exports = map[string]lua.LGFunction{
	"": newFn,
}

func newFn(L *lua.LState) int {
	client := FastWsClient{}
	ud := L.NewUserData()
	ud.Value = client
	L.SetMetatable(ud, L.GetTypeMetatable(SERVICE_TYPENAME))
	L.Push(ud)
	return 1

}

func registerClientType(L *lua.LState) {
	meta := L.NewTypeMetatable(SERVICE_TYPENAME)
	L.SetField(meta, "__index", L.SetFuncs(L.NewTable(), clientMethods))
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "_DEBUG", lua.LBool(false))
	L.SetField(mod, "_VERSION", lua.LString("1.0.0"))
	registerClientType(L)
	L.Push(mod)
	return 1
}
