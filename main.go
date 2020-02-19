package main

import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"github.com/lw000/gocommon/sys"
	log "github.com/sirupsen/logrus"
	"gogopherlua/luaapp"
	"os"
)

var (
	app *luaapp.LuaApp
)

func main() {
	tysys.RegisterOnInterrupt(func(sign os.Signal) {
		app.Stop()
		log.WithField("sign", fmt.Sprintf("%v", sign)).Error("GOLUA·退出")
	})

	app = luaapp.New()
	app.Start()
	app.Wait()

	// if err := svc.Run(&service.LuaService{}); err != nil {
	// 	log.Error(err)
	// }
}
