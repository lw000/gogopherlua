package export_wsclient

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func send(L *lua.LState) int {
	client := checkClient(L)

	if err := client.SendMessage([]byte{}); err != nil {
		log.Error(err)
		L.Push(lua.LNumber(0))
		return 1
	}
	L.Push(lua.LNumber(1))
	return 1
}
