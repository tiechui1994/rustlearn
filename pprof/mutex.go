package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func xxxx() {
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	for {
		stupidMutexUse()
	}
}

func stupidMutexUse() {
	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}
