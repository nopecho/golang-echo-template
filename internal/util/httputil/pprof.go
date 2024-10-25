package httputil

import (
	"github.com/rs/zerolog/log"
	"net/http"
	_ "net/http/pprof"
)

func ListenPprof() {
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Warn().Msgf("failed pprof server error: %v", err)
	}
}
