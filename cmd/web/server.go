package main

import (
	"github.com/nopecho/golang-template/internal/app/api"
	"github.com/nopecho/golang-template/internal/app/config"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/internal/app/svc"
	"github.com/nopecho/golang-template/pkg/echoserver"
	"github.com/nopecho/golang-template/pkg/gorm/datasource"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = datasource.NewPostgres(config.Env.Postgres.DSN(), datasource.DefaultConnPool())
}

func main() {
	server := echoserver.NewEcho()
	handler := api.NewHandler("v1")

	repository := database.NewDomainPostgresRepository(db)
	service := svc.NewDomainService(repository, nil)
	router := api.NewDomainRouter(service)

	handler.Register(router, router, router)
	handler.Route(server)
	echoserver.Run(server, config.Env.Port)
}
