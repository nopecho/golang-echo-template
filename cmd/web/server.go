package main

import (
	"github.com/nopecho/golang-template/internal/app/api/v1"
	"github.com/nopecho/golang-template/internal/app/config"
	"github.com/nopecho/golang-template/pkg/echoserver"
)

func main() {
	e := echoserver.NewEcho()

	apiV1 := echoserver.Version(e, 1)
	domainHandler := v1.NewDomainHandler(apiV1, nil)
	domainHandler.Routing()

	echoserver.Run(e, config.Env.Port)
}
