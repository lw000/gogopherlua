package export_ini

import (
	"github.com/Unknwon/goconfig"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"tostring":  toString,
	"tonumber":  toNumber,
	"toboolean": toBoolean,
}

func checkClient(L *lua.LState) *goconfig.ConfigFile {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*goconfig.ConfigFile); ok {
		return v
	}
	L.ArgError(1, "ini server expected")
	return nil
}

func toString(L *lua.LState) int {
	ini := checkClient(L)
	top := L.GetTop()
	if top == 2 {
		key := L.ToString(2)
		s, err := ini.GetValue("", key)
		if err != nil {
			log.Error(err)
			L.Push(lua.LString(err.Error()))
			L.Push(lua.LNumber(1))
			return 2
		}
		L.Push(lua.LString(s))
		return 1
	} else if top == 3 {
		section := L.ToString(2)
		key := L.ToString(3)
		s, err := ini.GetValue(section, key)
		if err != nil {
			log.Error(err)
			L.Push(lua.LString(err.Error()))
			L.Push(lua.LNumber(1))
			return 2
		}
		L.Push(lua.LString(s))
		return 1
	}

	L.Push(lua.LString("参数错误"))
	L.Push(lua.LNumber(1))
	return 2
}

func toNumber(L *lua.LState) int {
	ini := checkClient(L)
	top := L.GetTop()
	if top == 2 {
		key := L.ToString(2)
		s, err := ini.Int64("", key)
		if err != nil {
			log.Error(err)
			L.Push(lua.LString(err.Error()))
			L.Push(lua.LNumber(1))
			return 2
		}
		L.Push(lua.LNumber(s))
		return 1
	} else if top == 3 {
		section := L.ToString(2)
		key := L.ToString(3)
		s, err := ini.Int64(section, key)
		if err != nil {
			log.Error(err)
			L.Push(lua.LString(err.Error()))
			L.Push(lua.LNumber(1))
			return 2
		}
		L.Push(lua.LNumber(s))
		return 1
	}
	L.Push(lua.LString("参数错误"))
	L.Push(lua.LNumber(1))
	return 2
}

func toBoolean(L *lua.LState) int {
	ini := checkClient(L)
	top := L.GetTop()
	if top == 2 {
		key := L.ToString(2)
		s, err := ini.Bool("", key)
		if err != nil {
			log.Error(err)
			L.Push(lua.LString(err.Error()))
			L.Push(lua.LNumber(1))
			return 2
		}
		L.Push(lua.LBool(s))
		return 1
	} else if top == 3 {
		section := L.ToString(2)
		key := L.ToString(3)
		s, err := ini.Bool(section, key)
		if err != nil {
			log.Error(err)
			L.Push(lua.LString(err.Error()))
			L.Push(lua.LNumber(1))
			return 2
		}
		L.Push(lua.LBool(s))
		return 1
	}

	L.Push(lua.LString("参数错误"))
	L.Push(lua.LBool(false))
	return 2
}
