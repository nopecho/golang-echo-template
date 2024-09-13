package main

import (
	"github.com/joho/godotenv"
	"github.com/nopecho/golang-template/internal/app/api"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/internal/app/svc"
	"github.com/nopecho/golang-template/internal/pkg/helper"
	"github.com/nopecho/golang-template/pkg/echoserver"
	"github.com/nopecho/golang-template/pkg/gorm/datasource"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	helper.PrettyLogging()
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msgf("Error loading .env file: %v", err)
	}
}

func main() {
	dbInfo := database.EnvConnectionInfo()
	db = datasource.NewPostgres(dbInfo.DSN(), datasource.DefaultConnPool())
	db.AutoMigrate(&database.DomainEntity{}, &database.Domain2Entity{})

	repository := database.NewDomainPostgresRepository(db)
	service := svc.NewDomainService(repository, nil)
	router := api.NewDomainRouter(service)

	server := echoserver.NewEcho()

	handler := api.NewHandler("v1")
	handler.Register(router, router, router, router)
	handler.Route(server)

	handler2 := api.NewHandler("")
	handler2.Register(router)
	handler2.Route(server)

	echoserver.Run(server, helper.EnvPort())
}
