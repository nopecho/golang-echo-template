package main

import (
	"github.com/nopecho/golang-template/internal/app/api"
	"github.com/nopecho/golang-template/internal/app/config"
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/app/svc"
	"github.com/nopecho/golang-template/pkg/echoserver"
)

func init() {

}

func main() {
	e := echoserver.NewEcho()
	handler := api.NewHandler("v1")

	router := initRouter()
	handler.Register(router, router, router)

	handler.Route(e)
	echoserver.Run(e, config.Env.Port)
}

func initRouter() *api.DomainRouter {
	repository := domain.NewMemoryDomainRepository()
	service := svc.NewDomainService(repository, nil)
	router := api.NewDomainRouter(service)
	return router
}
