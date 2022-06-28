package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "net/http/pprof"
)

//go:noinline
func exec() []*int64 {
	s := make([]*int64, 204800, 204800)
	for i := 512; i < 768; i++ {
		s[i] = new(int64)
	}

	os.Open("xxx")

	return s[512:768]
}

// 运行一段时间：fatal error: runtime: out of memory
func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	var temp [][]*int64
	tick := time.Tick(200 * time.Millisecond)
	for range tick {
		value := exec()
		temp = append(temp, value)
	}
}
