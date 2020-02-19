package export_log

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

// log_error ...
func log_error(L *lua.LState) int {
	args := getLuaArgs(L)
	if len(args) > 0 {
		log.Error(args)
	}
	return 0
}
