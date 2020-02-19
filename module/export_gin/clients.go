package export_gin

import (
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"get":        get,
	"post":       post,
	"run":        run,
	"middleware": middleware,
}

func checkClient(L *lua.LState) *SafeApplication {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*SafeApplication); ok {
		return v
	}
	L.ArgError(1, "gin server expected")
	return nil
}
