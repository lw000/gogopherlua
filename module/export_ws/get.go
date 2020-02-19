package export_ws

import (
	"errors"
	"gogopherlua/helper"
	"gogopherlua/luacontext"
	"gogopherlua/pipe"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func addGET(L *lua.LState) int {
	app := checkClient(L)

	relativePath := L.ToString(2)
	if relativePath == "" {
		L.Push(lua.LNumber(0))
		return 1
	}

	cb := L.ToFunction(3)
	if cb == nil {
		L.Push(lua.LNumber(0))
		return 1
	}

	app.Engine().GET(relativePath, func(c *gin.Context) {
		args, err := url.ParseQuery(c.Request.URL.RawQuery)
		if err != nil {
			log.Error(err)
			return
		}

		t := L.NewTable()
		for k, v := range args {
			L.SetField(t, k, lua.LString(v[0]))
		}

		luaContext := luacontext.GetLuaContext(L)
		if luaContext == nil {
			err := errors.New("core error")
			log.Error(err)
			return
		}

		response := make(chan *pipe.Result, 1)
		luaContext.Pipe.Call(lua.P{Fn: cb, NRet: 1, Protect: true}, response, t)
		result := <-response
		close(response)
		if result.Len() > 0 {
			v := result.Get(0)
			if v.Type() == lua.LTTable {
				table := v.(*lua.LTable)
				m := helper.LuaTableToMap(table)
				c.JSON(http.StatusOK, m)
			}
		}
	})

	L.Push(lua.LNumber(1))
	return 1
}
