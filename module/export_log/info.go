package export_log

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

// log_info ...
func log_info(L *lua.LState) int {
	args := getLuaArgs(L)
	if len(args) > 0 {
		log.Info(args)
	}
	return 0
}
