package main

import (
	"github.com/nopecho/golang-template/internal/util/chore"
	"github.com/nopecho/golang-template/internal/util/echo"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go scheduling()
	e := echo.NewEcho()
	echo.Run(e, chore.EnvPort())
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
