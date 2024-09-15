package main

import (
	"github.com/nopecho/golang-template/internal/pkg/apputil"
	"github.com/nopecho/golang-template/internal/pkg/echoserver"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go scheduling()
	e := echoserver.NewEcho()
	echoserver.Run(e, apputil.EnvPort())
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
