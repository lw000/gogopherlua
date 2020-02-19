package export_cron

import (
	"github.com/robfig/cron"
	lua "github.com/yuin/gopher-lua"
)

const (
	SERVICE_TYPENAME = "corn{service}"
)

type SafeCorn struct {
	cr *cron.Cron
}

func newFn(L *lua.LState) int {
	cr := &SafeCorn{
		cr: cron.New(),
	}
	ud := L.NewUserData()
	ud.Value = cr
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
	L.PreloadModule("gcron", Loader)
}
