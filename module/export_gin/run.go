package export_gin

import (
	"github.com/lw000/gocommon/app/gin"
	"github.com/lw000/gocommon/web/gin/middleware"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func run(L *lua.LState) int {
	app := checkClient(L)
	port := L.ToInt64(2)
	go func() {
		err := app.app.Run(port, func(app *tygin.WebApplication) {
			// 跨域控制
			app.Engine().Use(
				tymiddleware.CorsHandler(nil),
			)
		})

		if err != nil {
			log.Error(err)
		}
	}()

	return 0
}
