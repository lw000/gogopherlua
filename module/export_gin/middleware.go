package export_gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/luacontext"
	"gogopherlua/pipe"
	"net/http"
)

func middleware(L *lua.LState) int {
	app := checkClient(L)

	cb := L.ToFunction(2)
	if cb == nil {
		L.Push(lua.LNumber(0))
		return 1
	}

	app.app.Engine().Use(func(c *gin.Context) {
		t := L.NewTable()
		L.SetField(t, "host", lua.LString(c.Request.Host))
		L.SetField(t, "ip", lua.LString(c.ClientIP()))

		header := L.NewTable()
		for k, v := range c.Request.Header {
			L.SetField(header, k, lua.LString(v[0]))
		}
		L.SetField(t, "header", header)

		luaContext := luacontext.GetLuaContext(L)
		if luaContext == nil {
			err := errors.New("core error")
			log.Error(err)
			return
		}

		response := make(chan *pipe.Result, 1)
		luaContext.Pipe.Call(lua.P{Fn: cb, NRet: 2, Protect: true}, response, t)
		result := <-response
		close(response)

		if result.Len() > 0 {
			v1 := result.Get(0)
			v2 := result.Get(1)

			var ok lua.LBool
			if v1.Type() == lua.LTBool {
				ok = v1.(lua.LBool)
			}

			var msg string
			if v2.Type() == lua.LTString {
				msg = string(v2.(lua.LString))
			}

			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"c": 0, "error": msg})
				return
			}
		}
		c.Next()
	})

	L.Push(lua.LNumber(1))
	return 1
}
