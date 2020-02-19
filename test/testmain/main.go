// test project main.go
package main

import (
	"gogopherlua/test/testmain/config"
	"log"
	"sync"
	"time"
)

var (
	successCount uint64
	failCount    uint64

	cfg *config.Config
	wg  *sync.WaitGroup
)

func main() {
	cfg = config.NewConfig()
	err := cfg.Load("./conf/conf.json")
	if err != nil {
		log.Println(err)
		return
	}

	wg = &sync.WaitGroup{}
	start := time.Now()
	if cfg.Method == "get" {
		DoGet(cfg)
	} else if cfg.Method == "post" {
		DoPost(cfg)
	}

	wg.Wait()

	end := time.Now()
	log.Printf("success:[%d], fail:[%d], [%v]", successCount, failCount, end.Sub(start))
}
