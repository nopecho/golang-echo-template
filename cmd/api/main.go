package main

import (
	"github.com/joho/godotenv"
	"github.com/nopecho/golang-template/internal/app/api"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/internal/app/svc"
	"github.com/nopecho/golang-template/internal/pkg/apputil"
	"github.com/nopecho/golang-template/internal/pkg/echoserver"
	"github.com/nopecho/golang-template/internal/pkg/gorm/datasource"
	"github.com/rs/zerolog/log"
)

func init() {
	apputil.PrettyLogging()
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msgf("Error loading .env file: %v", err)
	}
}

func main() {
	dbConn := datasource.DefaultConnectionInfo()
	db := datasource.NewPostgres(dbConn.DSN(), dbConn.ConnectionPool)
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

	echoserver.Run(server, apputil.EnvPort())
}
