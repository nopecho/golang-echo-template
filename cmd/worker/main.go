package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nopecho/golang-template/internal/util/common"
	"github.com/nopecho/golang-template/internal/util/echo"
	"github.com/nopecho/golang-template/internal/util/http"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go http.ListenPprof()
	go scheduling()
	e := echo.NewEcho()
	echo.Run(e, common.EnvPort())
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
