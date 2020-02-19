package export_cron

import (
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"add":   add,
	"start": start,
	"stop":  stop,
}

func checkClient(L *lua.LState) *SafeCorn {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*SafeCorn); ok {
		return v
	}
	L.ArgError(1, "SafeCorn server expected")
	return nil
}
