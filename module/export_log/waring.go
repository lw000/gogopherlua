package export_log

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

// log_waring ...
func log_waring(L *lua.LState) int {
	args := getLuaArgs(L)
	if len(args) > 0 {
		log.Warn(args)
	}
	return 0
}
