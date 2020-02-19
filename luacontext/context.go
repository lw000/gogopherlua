package luacontext

import (
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/pipe"
)

type LuaContext struct {
	id   int
	name string
	exit chan struct{}
	Pipe *pipe.LuaPipe
}

func (l *LuaContext) Exit() chan struct{} {
	return l.exit
}

func (l *LuaContext) Close() {
	if l.exit != nil {
		close(l.exit)
	}
}

func New(id int, name string) *LuaContext {
	return &LuaContext{
		id:   id,
		name: name,
		exit: make(chan struct{}),
		Pipe: pipe.New(),
	}
}

func GetLuaContext(L *lua.LState) *LuaContext {
	ud := L.GetGlobal("GO_LUA_CONTEXT")
	if ud.Type() == lua.LTUserData {
		if lv, ok := ud.(*lua.LUserData); ok {
			if v, ok := lv.Value.(*LuaContext); ok {
				return v
			}
		}
	}
	L.ArgError(1, "go lua context error")
	return nil
}
