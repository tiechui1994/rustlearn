package main

import (
	"fmt"
	"net/http"
	"os"

	_ "net/http/pprof"
)

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"aa":11}`))
	})
	mux.HandleFunc("/api/add", AddHandler)
	mux.HandleFunc("/api/get", GetHandler)
	mux.HandleFunc("/api/resize", ResizeHandler)
	mux.HandleFunc("/api/clear", ClearHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func main() {
	err := initCache(1024 * 1024)
	if err != nil {
		os.Exit(1)
	}

	go func() {
		ip := ":6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	startServer()
}
