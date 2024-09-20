package common

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PrettyLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
		NoColor:    true,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s", i))
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("[%s]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf(": %s", i)
		},
	}).With().Caller().Logger()
}
