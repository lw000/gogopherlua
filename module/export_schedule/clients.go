package export_schedule

import (
	"github.com/lw000/gocommon/schedule"
	"gogopherlua/helper"
	"gogopherlua/luacontext"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"start": start,
	"stop":  stop,
	"add":   add,
}

func checkClient(L *lua.LState) *SafeSchedule {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*SafeSchedule); ok {
		return v
	}
	L.ArgError(1, "SafeSchedule server expected")
	return nil
}

func start(L *lua.LState) int {
	sche := checkClient(L)
	if sche == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}
	if err := sche.sche.Start(); err != nil {
		L.Push(lua.LNumber(0))
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LNumber(1))
	return 1
}

func stop(L *lua.LState) int {
	sche := checkClient(L)
	if sche == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}
	sche.sche.Stop()
	return 0
}

func add(L *lua.LState) int {
	sche := checkClient(L)
	if sche == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}

	s := L.ToInt64(2)

	cb := L.ToFunction(3)
	if cb == nil {
		L.Push(lua.LNumber(0))
		L.Push(lua.LString("回调函数为空"))
		return 2
	}

	args := L.ToTable(4)
	data := helper.LuaTableToMap(args)

	id := sche.sche.AddTask(s, &tyschedule.Task{Data: data, Fn: func(data interface{}) {
		t := L.NewTable()
		for k, v := range data.(map[string]interface{}) {
			L.SetField(t, k, v.(lua.LString))
		}
		luaContext := luacontext.GetLuaContext(L)
		if luaContext == nil {
			log.Error("core error")
			return
		}
		luaContext.Pipe.Call(lua.P{Fn: cb, NRet: 0, Protect: true}, nil, t)
	}})
	if id == "" {
		L.Push(lua.LNumber(0))
		L.Push(lua.LString("添加任务失败"))
		return 2
	}

	L.Push(lua.LNumber(1))
	L.Push(lua.LString(id))
	return 2
}
