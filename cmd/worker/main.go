package main

import (
	"github.com/nopecho/golang-template/internal/utils/chore"
	"github.com/nopecho/golang-template/internal/utils/echoutils"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go scheduling()
	e := echoutils.NewEcho()
	echoutils.Run(e, chore.EnvPort())
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
