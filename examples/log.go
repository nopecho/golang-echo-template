package examples

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func setupTestLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02T15:04:05.000",
	})
}
