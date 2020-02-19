package export_ws

import (
	"github.com/lw000/gocommon/app/gin"
	"github.com/lw000/gocommon/web/gin/middleware"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func runGin(L *lua.LState) int {
	app := checkClient(L)
	port := L.ToInt64(2)
	er := app.Run(port, func(app *tygin.WebApplication) {
		// 跨域控制
		app.Engine().Use(
			tymiddleware.CorsHandler(nil),
		)
	})

	if er != nil {
		log.Error(er)
	}

	return 0
}
