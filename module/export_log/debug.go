package export_log

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

// log_debug ...
func log_debug(L *lua.LState) int {
	args := getLuaArgs(L)
	if len(args) > 0 {
		log.Debug(args)
	}
	return 0
}
