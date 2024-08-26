package app

import (
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	handlers "github.com/billowdev/exclusive-go-hexa/internal/adapters/http/handlers/system_fields"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/http/routers"
	repositories "github.com/billowdev/exclusive-go-hexa/internal/adapters/repositories/system_fields"
	services "github.com/billowdev/exclusive-go-hexa/internal/core/services/system_fields"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	SystemFieldApp(route, db)
	return app
}

func SystemFieldApp(r routers.RouterImpls, db *gorm.DB) {
	transactorRepo := database.NewTransactorRepo(db)
	sfRepo := repositories.NewSystemFieldRepo(db)
	sfSrv := services.NewSystemFieldService(sfRepo, transactorRepo)
	sfHandlers := handlers.NewSystemFieldHandler(sfSrv)
	r.CreateSystemFieldRoute(sfHandlers)
}
