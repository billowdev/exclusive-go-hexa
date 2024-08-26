package app

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database"
	repositories "github.com/billowdev/document-system-field-manager/internal/adapters/repositories/system_fields"
	handlers "github.com/billowdev/document-system-field-manager/internal/adapters/rest/handlers/system_fields"
	"github.com/billowdev/document-system-field-manager/internal/adapters/rest/routers"
	services "github.com/billowdev/document-system-field-manager/internal/core/services/system_fields"
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
