package common

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func SetUpApplication() {
	go listenPprof()
	prettyLogging()
	logSystemInfo()
}

func listenPprof() {
	_ = http.ListenAndServe(":6060", nil)
}

func logSystemInfo() {
	log.Info().
		Str("GOLANG", runtime.Version()).
		Str("GOARCH", runtime.GOARCH).
		Str("GOOS", runtime.GOOS).
		Int("NUM_CPU", runtime.NumCPU()).
		Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
		Msg("runtime")
}

func prettyLogging() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: time.RFC3339Nano,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s [", i))
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s ]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf(": %s", i)
		},
	}).With().Caller().Logger()
}
