package export_log

import (
	"bytes"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/helper"
)

var exports = map[string]lua.LGFunction{
	"debug":  log_debug,
	"info":   log_info,
	"waring": log_waring,
	"error":  log_error,
}

func getLuaArgs(L *lua.LState) string {
	top := L.GetTop()
	if top > 0 {
		buf := bytes.Buffer{}
		for i := 1; i <= top; i++ {
			lv := L.Get(i)
			switch lv.Type() {
			case lua.LTTable:
				s := helper.PrintLuaTableV2(lv.(*lua.LTable))
				buf.WriteString(s)
				buf.WriteString(" ")
			default:
				buf.WriteString(lv.String())
				buf.WriteString(" ")
			}
		}
		return buf.String()
	}
	return ""
}

// Loader 加载库函数
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "_VERSION", lua.LString("1.0.0"))
	L.Push(mod)
	return 1
}

func Preload(L *lua.LState) {
	L.PreloadModule("glog", Loader)
}
