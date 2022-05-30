package main

import (
	"net/http"
	"os"
)

func startServer() {
	mux := http.NewServeMux()
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

	startServer()
}
