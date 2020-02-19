package locker

import "sync"

var GlobalMutex *sync.Mutex

func init() {
	GlobalMutex = &sync.Mutex{}
}
