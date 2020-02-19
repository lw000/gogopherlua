package service

import (
	"errors"
	"github.com/judwhite/go-svc/svc"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"gogopherlua/module"
	"gogopherlua/service/module"
)

type LuaService struct {
	servs   []*module.LuaServiceModule
	modules []string
}

func (s *LuaService) Init(env svc.Environment) error {
	var err error
	s.modules, err = s.getModules()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *LuaService) Start() error {
	for _, mod := range s.modules {
		serv := module.New(mod)
		err := serv.Start()
		if err != nil {
			log.Error(err)
			continue
		}
		s.servs = append(s.servs, serv)
	}
	return nil
}

func (s *LuaService) getModules() ([]string, error) {
	L := lua.NewState()
	defer L.Close()

	export_module.Preload(L)

	var err error
	if err = L.DoFile("./script/main.lua"); err != nil {
		log.Error(err)
		return nil, err
	}

	modulePaths := L.GetGlobal("__TY_CHILD_MODULES_PATH__")
	switch modulePaths.Type() {
	case lua.LTNil:
		log.Println("__TY_CHILD_MODULES_PATH__ not found")
		return nil, err
	case lua.LTTable:
		var modules []string
		t := modulePaths.(*lua.LTable)
		t.ForEach(func(i lua.LValue, value lua.LValue) {
			modName := value.String()
			if modName != "" {
				modules = append(modules, modName)
			}
		})
		return modules, nil
	}
	return nil, errors.New("未知错误")
}

func (s *LuaService) Stop() error {
	for _, srv := range s.servs {
		srv.Stop()
	}
	log.Error("LuaService exit")
	return nil
}
