package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nopecho/golang-template/internal/app/api"
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/internal/util/common"
	"github.com/nopecho/golang-template/internal/util/echoutil"
	"github.com/nopecho/golang-template/internal/util/gormutil/datasource"
)

func init() {
	common.SetUpApplication()
}

func main() {
	db := datasource.NewDefaultPostgres()
	_ = db.AutoMigrate(&database.AnyEntity{}, &database.Any2Entity{})

	handler := api.NewVersionHandler("v1")
	var (
		repository = database.NewAnyGormRepository(db)
		service    = domain.NewAnyService(repository)
		router     = api.NewAnyRouter(service)
	)
	handler.Register(router)

	handler2 := api.NewRootHandler()
	handler2.Register(router)

	server := echoutil.NewEcho()
	handler.Routing(server)
	handler2.Routing(server)

	ctx := context.Background()
	echoutil.Run(ctx, server, common.EnvPort())
}
