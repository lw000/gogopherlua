package export_loop

import (
	"context"

	lua "github.com/yuin/gopher-lua"
)

const (
	SERVICE_TYPENAME = "runloop{service}"
)

type RunLoop struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func newFn(L *lua.LState) int {
	rl := &RunLoop{}
	rl.ctx, rl.cancel = context.WithCancel(context.Background())

	ud := L.NewUserData()
	ud.Value = rl
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
