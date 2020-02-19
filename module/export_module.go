package export_module

import (
	"gogopherlua/module/export_cron"
	"gogopherlua/module/export_function"
	"gogopherlua/module/export_gin"
	"gogopherlua/module/export_ini"
	"gogopherlua/module/export_log"
	"gogopherlua/module/export_loop"
	"gogopherlua/module/export_mysql"
	"gogopherlua/module/export_redis"
	"gogopherlua/module/export_schedule"
	"gogopherlua/module/export_ws"
	"gogopherlua/module/export_wsclient"
	json "layeh.com/gopher-json"
	"net/http"

	"github.com/cjoudrey/gluahttp"
	"github.com/cjoudrey/gluaurl"
	"github.com/nubix-io/gluasocket"
	"github.com/tengattack/gluacrypto"
	"github.com/tengattack/gluasql"
	lua "github.com/yuin/gopher-lua"
)

// Loader 加载库函数
func Preload(L *lua.LState) int {
	export_log.Preload(L)
	export_gin.Preload(L)
	export_ini.Preload(L)
	export_schedule.Preload(L)
	export_cron.Preload(L)

	L.PreloadModule("gws", export_ws.Loader)
	L.PreloadModule("gwsclient", export_wsclient.Loader)
	L.PreloadModule("gloop", export_loop.Loader)
	L.PreloadModule("gmysql", export_mysql.Loader)

	L.PreloadModule("ghttp", gluahttp.NewHttpModule(&http.Client{}).Loader)
	L.PreloadModule("gurl", gluaurl.Loader)

	export_function.Loader(L)
	export_redis.Preload(L)

	json.Preload(L)
	gluasql.Preload(L)
	gluasocket.Preload(L)
	gluacrypto.Preload(L)

	return 1
}
