package export_function

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/luacontext"
	"time"
)

// loop ...
func loop(L *lua.LState) int {
	top := L.GetTop()
	if top >= 1 {
		s := L.ToInt64(1)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(s))
		select {
		case <-ctx.Done():
			cancel()
			return 0
		}
	} else {
		ctx, cancel := context.WithCancel(context.Background())
		select {
		case <-ctx.Done():
			cancel()
			return 0
		}
	}
}

func sleep(L *lua.LState) int {
	top := L.GetTop()
	if top >= 1 {
		s := L.ToInt64(1)
		time.Sleep(time.Millisecond * time.Duration(s))
	}
	return 0
}

func after(L *lua.LState) int {
	top := L.GetTop()
	if top >= 2 {
		s := L.ToInt64(1)
		cb := L.ToFunction(2)

		luaContext := luacontext.GetLuaContext(L)
		if luaContext == nil {
			err := errors.New("core error")
			log.Error(err)
			return 0
		}

		time.AfterFunc(time.Millisecond*time.Duration(s), func() {
			luaContext.Pipe.Call(lua.P{Fn: cb, NRet: 0, Protect: true}, nil)
		})
	}

	return 0
}

// Loader 加载库函数
func Loader(L *lua.LState) int {
	L.SetGlobal("loop", L.NewFunction(loop))
	L.SetGlobal("sleep", L.NewFunction(sleep))
	L.SetGlobal("after", L.NewFunction(after))
	return 1
}
