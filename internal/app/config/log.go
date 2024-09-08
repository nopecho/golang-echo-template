package config

import (
	"github.com/nopecho/golang-template/internal/pkg/dayte"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func PrettyLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: dayte.DateTime,
	})
}
