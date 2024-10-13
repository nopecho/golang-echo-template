package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nopecho/golang-template/internal/app/api"
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/internal/util/common"
	"github.com/nopecho/golang-template/internal/util/echo"
	"github.com/nopecho/golang-template/internal/util/gorm/datasource"
)

func init() {
	common.SetUpApplication()
}

func main() {
	db := datasource.NewDefaultPostgres()
	_ = db.AutoMigrate(&database.AnyEntity{}, &database.Any2Entity{})

	var (
		repository = database.NewAnyGormRepository(db)
		service    = domain.NewAnyService(repository)
		router     = api.NewAnyRouter(service)
		handler    = api.NewVersionHandler("v1")
	)
	handler.Register(router, router, router, router)

	handler2 := api.NewHandler()
	handler2.Register(router)

	server := echo.NewEcho()
	handler.Routing(server)
	handler2.Routing(server)

	echo.Run(server, common.EnvPort())
}
