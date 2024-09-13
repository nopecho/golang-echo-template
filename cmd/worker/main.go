package main

import (
	"github.com/nopecho/golang-template/internal/pkg/helper"
	"github.com/nopecho/golang-template/pkg/echoserver"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go scheduling()
	e := echoserver.NewEcho()
	echoserver.Run(e, helper.EnvPort())
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
