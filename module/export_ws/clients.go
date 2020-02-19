package export_ws

import (
	"github.com/lw000/gocommon/app/gin"

	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"get":  addGET,
	"post": addPOST,
	"run":  runGin,
}

func checkClient(L *lua.LState) *tygin.WebApplication {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*tygin.WebApplication); ok {
		return v
	}
	L.ArgError(1, "gin server expected")
	return nil
}
