package export_ini

import (
	"github.com/Unknwon/goconfig"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

const (
	SERVICE_TYPENAME = "ini{service}"
)

var exports = map[string]lua.LGFunction{
	"new": newFn,
}

func newFn(L *lua.LState) int {
	file := L.ToString(1)
	f, err := goconfig.LoadConfigFile(file)
	if err != nil {
		log.Error(err)

		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	ud := L.NewUserData()
	ud.Value = f
	L.SetMetatable(ud, L.GetTypeMetatable(SERVICE_TYPENAME))

	L.Push(ud)
	return 1
}

func registerClientType(L *lua.LState) {
	mt := L.NewTypeMetatable(SERVICE_TYPENAME)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), clientMethods))
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "_DEBUG", lua.LBool(true))
	L.SetField(mod, "_VERSION", lua.LString("1.0.0"))

	registerClientType(L)

	L.Push(mod)
	return 1
}

func Preload(L *lua.LState) {
	L.PreloadModule("gini", Loader)
}
