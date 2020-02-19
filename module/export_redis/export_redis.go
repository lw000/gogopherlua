package export_redis

import (
	"github.com/lw000/gocommon/db/rdsex"

	lua "github.com/yuin/gopher-lua"
)

const (
	SERVICE_TYPENAME = "redis{service}"
)

func newFn(L *lua.LState) int {
	cfgpath := L.ToString(1)
	cfg := &tyrdsex.JsonConfig{}
	err := cfg.Load(cfgpath)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	rds := &tyrdsex.RdsServer{}
	err = rds.OpenWithJsonConfig(cfg)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	ud := L.NewUserData()
	ud.Value = rds
	L.SetMetatable(ud, L.GetTypeMetatable(SERVICE_TYPENAME))
	L.Push(ud)
	return 1
}

var exports = map[string]lua.LGFunction{
	"new": newFn,
}

func registerClientType(L *lua.LState) {
	mt := L.NewTypeMetatable(SERVICE_TYPENAME)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), clientMethods))
}

// Loader 加载库函数
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "_DEBUG", lua.LBool(false))
	L.SetField(mod, "_VERSION", lua.LString("1.0.0"))

	registerClientType(L)

	L.Push(mod)
	return 1
}

func Preload(L *lua.LState) {
	L.PreloadModule("gredis", Loader)
}
