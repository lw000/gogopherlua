package main

import (
	"gogopherlua/test/network"
	"gogopherlua/test/testmain/config"
	"net/url"
	"time"
)

// DoGet GET请求
func DoGet(cfg *config.Config) {
	u := url.Values{}
	u.Add("a", "11111111111111")
	u.Add("b", "222222222222222")

	if cfg.TLS {
		var i int64
		for i = 0; i < cfg.Count; i++ {
			for _, a := range cfg.HTTPS.Get {
				go network.HttpGet(wg, cfg.TLS, a+"?"+u.Encode())
				if cfg.Millisecond != 0 {
					time.Sleep(time.Millisecond * time.Duration(cfg.Millisecond))
				}
			}
		}
	} else {
		var i int64
		for i = 0; i < cfg.Count; i++ {
			for _, a := range cfg.HTTP.Get {
				go network.HttpGet(wg, cfg.TLS, a+"?"+u.Encode())
				if cfg.Millisecond != 0 {
					time.Sleep(time.Millisecond * time.Duration(cfg.Millisecond))
				}
			}
		}
	}
}

// DoPost POST请求
func DoPost(cfg *config.Config) {
	u := url.Values{}
	u.Add("a", "11111111111111")
	u.Add("b", "222222222222222")

	if cfg.TLS {
		var i int64
		for i = 0; i < cfg.Count; i++ {
			for _, a := range cfg.HTTPS.Post {
				go network.HttpPost(wg, cfg.TLS, a, nil, u)
				if cfg.Millisecond != 0 {
					time.Sleep(time.Millisecond * time.Duration(cfg.Millisecond))
				}
			}
		}
	} else {
		var i int64
		for i = 0; i < cfg.Count; i++ {
			for _, a := range cfg.HTTP.Post {
				go network.HttpPost(wg, cfg.TLS, a, nil, u)
				if cfg.Millisecond != 0 {
					time.Sleep(time.Millisecond * time.Duration(cfg.Millisecond))
				}
			}
		}
	}
}
