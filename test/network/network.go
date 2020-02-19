package network

import (
	"github.com/lw000/gocommon/network/http"
	"log"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

var defaultTimeout time.Duration = 15
var successCount uint64
var failCount uint64

func HttpPost(w *sync.WaitGroup, tls bool, url string, headers map[string]string, data url.Values) {
	if w != nil {
		w.Add(1)
		defer func() {
			defer w.Done()
		}()
	}

	var (
		er   error
		body []byte
	)

	if tls {
		_, body, er = tyhttp.DoHttpsPost(url, headers, data.Encode(), time.Second*defaultTimeout)
	} else {
		_, body, er = tyhttp.DoHttpPost(url, headers, data.Encode(), time.Second*defaultTimeout)
	}

	if er != nil {
		log.Println(atomic.AddUint64(&failCount, 1), er.Error())
		return
	}

	log.Printf("[%d] %s", atomic.AddUint64(&successCount, 1), string(body))
}

func HttpGet(w *sync.WaitGroup, tls bool, url string) {
	if w != nil {
		w.Add(1)
		defer func() {
			defer w.Done()
		}()
	}

	var (
		er   error
		body []byte
	)

	if tls {
		_, body, er = tyhttp.DoHttpsGet(url, nil, time.Second*defaultTimeout)
	} else {
		_, body, er = tyhttp.DoHttpGet(url, nil, time.Second*defaultTimeout)
	}

	if er != nil {
		log.Println(atomic.AddUint64(&failCount, 1), string(body), er.Error())
		return
	}
	log.Printf("[%d] %s", atomic.AddUint64(&successCount, 1), string(body))
}
