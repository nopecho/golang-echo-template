package main

import (
	"github.com/nopecho/golang-template/internal/app/config"
	"github.com/nopecho/golang-template/pkg/bootstrap/http"
)

func main() {
	c := config.EnvConfig
	server := http.NewHttpServer(&http.Option{
		Port:    c.Port,
		Handler: http.NewCompositeHandler(),
	})
	server.Run()
}
