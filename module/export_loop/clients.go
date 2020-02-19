package export_loop

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"loop":   loop,
	"cancel": cancel,
}

func checkClient(L *lua.LState) *RunLoop {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*RunLoop); ok {
		return v
	}
	L.ArgError(1, "RunLoop server expected")
	return nil
}

func loop(L *lua.LState) int {
	rl := checkClient(L)
	if rl == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}

	select {
	case <-rl.ctx.Done():
		break
	}

	L.Push(lua.LNumber(1))
	return 1
}

func cancel(L *lua.LState) int {
	rl := checkClient(L)
	if rl == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}

	rl.cancel()

	L.Push(lua.LNumber(1))
	return 1
}
