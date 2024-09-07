package main

import (
	"fmt"
	"github.com/nopecho/golang-template/internal/app/config"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})

	http.ListenAndServe(fmt.Sprintf(":%d", config.EnvConfig.Port), mux)
}
