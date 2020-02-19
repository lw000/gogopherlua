package helper

import (
	"bytes"
	lua "github.com/yuin/gopher-lua"
	"strings"
)

func LuaTableToMap(table *lua.LTable) map[string]interface{} {
	if table == nil {
		return nil
	}

	if table.Type() != lua.LTTable {
		return nil
	}

	result := make(map[string]interface{})
	table.ForEach(func(k lua.LValue, v lua.LValue) {
		switch v.Type() {
		case lua.LTTable:
			result[k.String()] = LuaTableToMap(v.(*lua.LTable))
		default:
			result[k.String()] = v.String()
		}
	})
	return result
}

func PrintLuaTableV2(table *lua.LTable) string {
	if table == nil {
		return ""
	}

	if table.Type() != lua.LTTable {
		return ""
	}

	buf := bytes.Buffer{}
	buf.WriteString("{")
	table.ForEach(func(k lua.LValue, v lua.LValue) {
		switch v.Type() {
		case lua.LTTable:
			buf.WriteString(k.String())
			buf.WriteString("=")
			s := PrintLuaTableV2(v.(*lua.LTable))
			if s != "" {
				buf.WriteString(s)
				buf.WriteString(",")
			}
		default:
			buf.WriteString(k.String())
			buf.WriteString("=")
			buf.WriteString(v.String())
			buf.WriteString(",")
		}
	})
	buf.WriteString("}")
	s := buf.String()
	s = strings.Replace(s, ",}", "}", 1)
	return s
}
