package export_redis

import (
	"encoding/json"
	"github.com/lw000/gocommon/db/rdsex"
	"gogopherlua/helper"

	lua "github.com/yuin/gopher-lua"
)

var clientMethods = map[string]lua.LGFunction{
	"close": redis_close,
	"get":   redis_get,
	"set":   redis_set,
}

func checkClient(L *lua.LState) *tyrdsex.RdsServer {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*tyrdsex.RdsServer); ok {
		return v
	}
	L.ArgError(1, "gredis server expected")
	return nil
}

func redis_close(L *lua.LState) int {
	rds := checkClient(L)
	err := rds.Close()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}

func redis_set(L *lua.LState) int {
	rds := checkClient(L)
	top := L.GetTop()
	switch top {
	case 3:
		var (
			s   string
			err error
		)
		key := L.ToString(2)
		typeValue := L.Get(3)
		switch typeValue.Type() {
		case lua.LTBool:
			value := L.ToBool(3)
			s, err = rds.Set(key, bool(value), -1)
		case lua.LTString:
			value := L.ToString(3)
			s, err = rds.Set(key, value, -1)
		case lua.LTNumber:
			value := L.ToNumber(3)
			s, err = rds.Set(key, float64(value), -1)
		case lua.LTTable:
			value := L.ToTable(3)
			m := helper.LuaTableToMap(value)
			if len(m) > 0 {
				var data []byte
				data, err = json.Marshal(m)
				if err != nil {
					L.Push(lua.LNil)
					L.Push(lua.LString(err.Error()))
					return 2
				}
				s, err = rds.Set(key, string(data), -1)
			}
		default:
		}
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(s))
		L.Push(lua.LNil)
		return 2
	}
	L.Push(lua.LNil)
	L.Push(lua.LString("参数错误"))
	return 2
}

func redis_get(L *lua.LState) int {
	rds := checkClient(L)
	top := L.GetTop()
	switch top {
	case 2:
		key := L.ToString(2)
		s, err := rds.Get(key)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(s))
		L.Push(lua.LNil)
		return 2
	}
	L.Push(lua.LNil)
	L.Push(lua.LString("参数错误"))
	return 2
}
