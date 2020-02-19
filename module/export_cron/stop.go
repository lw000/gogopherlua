package export_cron

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func stop(L *lua.LState) int {
	cron := checkClient(L)
	if cron == nil {
		log.Error("error object")
		L.Push(lua.LNumber(0))
		return 1
	}

	cron.cr.Stop()

	L.Push(lua.LNumber(1))
	return 1
}
