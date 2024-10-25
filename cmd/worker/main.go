package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nopecho/golang-template/internal/util/common"
	"github.com/nopecho/golang-template/internal/util/echoutil"
	"github.com/nopecho/golang-template/internal/util/httputil"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	go httputil.ListenPprof()
	go scheduling()
	e := echoutil.NewEcho()
	echoutil.Run(e, common.EnvPort())
}

func scheduling() {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		log.Info().Msgf("Tick: %v", t)
	}
}
