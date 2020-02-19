package module

import (
	"errors"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/global"
	"gogopherlua/module"
	"sync"
)

type LuaServiceModule struct {
	moduleName string
	exit       chan struct{}
	wg         sync.WaitGroup
}

func New(moduleName string) *LuaServiceModule {
	return &LuaServiceModule{
		exit:       make(chan struct{}, 1),
		moduleName: moduleName,
	}
}

func (s *LuaServiceModule) Start() error {
	if s.moduleName == "" {
		return errors.New("moduleName is empty")
	}

	s.wg.Add(1)
	go s.run()

	return nil
}

func (s *LuaServiceModule) run() {
	defer func() {
		if x := recover(); x != nil {
			log.Error(x)
		}
		log.WithField("module", s.moduleName).Error("子服务模块·退出")
		s.wg.Done()
	}()
	log.WithField("module", s.moduleName).Info("子服务模块·启动")

	// opt := lua.Options{
	// 	RegistrySize:     1024 * 20, // this is the initial size of the registry
	// 	RegistryMaxSize:  1024 * 80, // this is the maximum size that the registry can grow to. If set to `0` (the default) then the registry will not auto grow
	// 	RegistryGrowStep: 32,        // this is how much to step up the registry by each time it runs out of space. The default is `32`.
	// }
	// L := lua.NewState(opt)

	L := lua.NewState()
	defer L.Close()

	export_module.Preload(L)

	if err := L.DoFile(s.moduleName); err != nil {
		log.Error(err)
		return
	}

	for {
		select {
		case item := <-global.LuaCommonPipe.Wait():
			var result global.Result
			err := L.CallByParam(item.Request.P, item.Request.Args...)
			if err != nil {
				log.Error(err)
				if item.Reponse != nil {
					item.Reponse <- &result
				}
			} else {
				if item.Request.P.NRet > 0 {
					for i := -item.Request.P.NRet; i < 0; i++ {
						v := L.Get(i)
						result.Add(v)
					}
					L.Pop(item.Request.P.NRet)
				}
				if item.Reponse != nil {
					item.Reponse <- &result
				}
			}
		case <-s.exit:
			return
		}
	}
}

func (s *LuaServiceModule) Stop() {
	close(s.exit)
	s.wg.Wait()
}
