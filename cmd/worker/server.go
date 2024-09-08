package main

import (
	"github.com/nopecho/golang-template/internal/app/config"
	"github.com/nopecho/golang-template/pkg/echoserver"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go scheduling()
	e := echoserver.NewEcho()
	echoserver.Run(e, config.Env.Port)
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
