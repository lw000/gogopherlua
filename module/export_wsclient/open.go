package export_wsclient

import (
	"errors"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/luacontext"
)

func open(L *lua.LState) int {
	top := L.GetTop()
	if top < 7 {
		L.Push(lua.LString("参数错误"))
		L.Push(lua.LNumber(0))
		return 2
	}

	client := checkClient(L)
	scheme := L.ToString(2)
	host := L.ToString(3)
	path := L.ToString(4)
	onConnected := L.ToFunction(5)
	onDisConnected := L.ToFunction(6)
	onMessage := L.ToFunction(7)

	if err := client.Open(scheme, host, path); err != nil {
		log.Error(err)
		L.Push(lua.LString(err.Error()))
		L.Push(lua.LNumber(0))
		return 2
	}

	luaContext := luacontext.GetLuaContext(L)
	if luaContext == nil {
		err := errors.New("core error")
		log.Error(err)
		L.Push(lua.LString(err.Error()))
		L.Push(lua.LNumber(0))
		return 2
	}

	client.HandleConnected(func() {
		luaContext.Pipe.Call(lua.P{Fn: onConnected, NRet: 0, Protect: true}, nil, lua.LString("ws连接成功"))
	})

	client.HandleDisConnected(func() {
		luaContext.Pipe.Call(lua.P{Fn: onDisConnected, NRet: 0, Protect: true}, nil, lua.LString("ws连接成功"))
	})

	client.HandleMessage(func(data []byte) {
		t := L.NewTable()
		for i, v := range data {
			t.RawSet(lua.LNumber(i), lua.LNumber(v))
		}
		luaContext.Pipe.Call(lua.P{Fn: onMessage, NRet: 0, Protect: true}, nil, t)
	})

	L.Push(lua.LNumber(1))
	return 1
}
