package export_cron

import (
	"errors"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/luacontext"
)

func add(L *lua.LState) int {
	cron := checkClient(L)
	if cron == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}

	spec := L.ToString(2)
	if spec == "" {
		log.Error("spec is empty")
		L.Push(lua.LNumber(0))
		return 1
	}

	cb := L.ToFunction(3)
	if cb == nil {
		log.Error("cb is empty")
		L.Push(lua.LNumber(0))
		return 1
	}

	luaContext := luacontext.GetLuaContext(L)
	if luaContext == nil {
		err := errors.New("core error")
		log.Error(err)
		L.Push(lua.LNumber(0))
		return 1
	}

	err := cron.cr.AddFunc(spec, func() {
		luaContext.Pipe.Call(lua.P{Fn: cb, NRet: 0, Protect: true}, nil)
	})

	if err != nil {
		log.Error(err)
		L.Push(lua.LNumber(0))
		return 1
	}

	L.Push(lua.LNumber(1))
	return 1
}
