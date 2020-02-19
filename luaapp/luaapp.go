package luaapp

import (
	"errors"
	"gogopherlua/luacontext"
	"gogopherlua/module"
	"gogopherlua/pipe"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"golang.org/x/sync/errgroup"
)

type LuaApp struct {
	sync.Mutex
	g       errgroup.Group
	ctxs    []*luacontext.LuaContext
	modules []string
	conf    map[string]string
}

func New() *LuaApp {
	return &LuaApp{
		conf: make(map[string]string),
	}
}

func (app *LuaApp) Start() {
	err := app.parseConfig()
	if err != nil {
		log.Error(err)
		return
	}

	configLocalFilesystemLogger("log", "tgolua", time.Hour*24*365, time.Hour*24)

	for k, v := range app.conf {
		switch k {
		case "debug":
			if v == "0" {
				log.SetLevel(log.ErrorLevel)
			} else {
				log.SetLevel(log.DebugLevel)
			}
		}
	}

	for i, mod := range app.modules {
		iCopy := i
		modCopy := mod
		app.Go(func() error {
			defer func() {
				if x := recover(); x != nil {
					log.Error(x)
				}
				log.WithField("module", modCopy).Info("子服务模块·退出")
			}()

			log.WithField("module", modCopy).Info("子服务模块·启动")

			// opt := lua.Options{
			// 	RegistrySize:     1024 * 20, // this is the initial size of the registry
			// 	RegistryMaxSize:  1024 * 80, // this is the maximum size that the registry can grow to. If set to `0` (the default) then the registry will not auto grow
			// 	RegistryGrowStep: 32,        // this is how much to step up the registry by each time it runs out of space. The default is `32`.
			// }
			// L := lua.NewState(opt)

			L := lua.NewState()
			defer L.Close()

			export_module.Preload(L)

			luaCtx := luacontext.New(iCopy, modCopy)
			ud := L.NewUserData()
			ud.Value = luaCtx
			L.SetMetatable(ud, L.GetTypeMetatable("GO_LUA_CONTEXT"))
			L.SetGlobal("GO_LUA_CONTEXT", ud)

			if err = L.DoFile(modCopy); err != nil {
				log.Error(err)
				return err
			}

			app.register(luaCtx)

		__EXIT:
			for {
				select {
				case data := <-luaCtx.Pipe.Wait():
					request := data.Request()
					err = L.CallByParam(request.P(), request.Args()...)
					if err != nil {
						log.Error(err)
						data.SetResponse(&pipe.Result{})
						break
					}

					var result pipe.Result
					if request.NRet() > 0 {
						for i := -request.NRet(); i < 0; i++ {
							v := L.Get(i)
							result.Add(v)
						}
						L.Pop(request.NRet())
					}
					data.SetResponse(&result)
				case <-luaCtx.Exit():
					goto __EXIT
				}
			}
		})
	}
}

func (app *LuaApp) register(ctx *luacontext.LuaContext) {
	app.Lock()
	defer app.Unlock()
	app.ctxs = append(app.ctxs, ctx)
}

func (app *LuaApp) parseConfig() error {
	L := lua.NewState()
	defer L.Close()

	export_module.Preload(L)

	var err error
	if err = L.DoFile("script/config.lua"); err != nil {
		log.Error(err)
		return err
	}

	tyModulePaths := L.GetGlobal("__TY_CHILD_MODULES_PATH__")
	switch tyModulePaths.Type() {
	case lua.LTNil:
		log.Error("__TY_CHILD_MODULES_PATH__ not found")
		return err
	case lua.LTTable:
		t := tyModulePaths.(*lua.LTable)
		t.ForEach(func(key lua.LValue, value lua.LValue) {
			s := value.String()
			if s != "" {
				app.modules = append(app.modules, s)
			}
		})
	default:
		return errors.New("模块[__TY_CHILD_MODULES_PATH__]错误")
	}

	tyConfigs := L.GetGlobal("__TY_CONFIG__")
	switch tyConfigs.Type() {
	case lua.LTNil:
		log.Error("__TY_CONFIG__ not found")
		return err
	case lua.LTTable:
		t := tyConfigs.(*lua.LTable)
		t.ForEach(func(key lua.LValue, value lua.LValue) {
			app.conf[key.String()] = value.String()
		})
	default:
		return errors.New("配置[__TY_CONFIG__]错误")
	}

	return nil
}

func (app *LuaApp) Stop() {
	for _, ctx := range app.ctxs {
		ctx.Close()
	}
}

func (app *LuaApp) Go(fn func() error) {
	app.g.Go(fn)
}

func (app *LuaApp) Wait() {
	if err := app.g.Wait(); err != nil {
		log.Error(err)
	}
	log.Error("LuaApp·exit")
}
