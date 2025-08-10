package http

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/di"
	"github.com/bagusyanuar/app-inventory-be/internal/http/middleware"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerDI) {
	app := cfg.App

	auth := app.Group("/auth")
	auth.Post("/login", handler.Auth.Login)

	jwtMiddleware := middleware.VerifyJWT(cfg)

	unit := app.Group("/unit", jwtMiddleware)
	unit.Post("/", handler.Unit.Create)
	unit.Get("/", handler.Unit.FindAll)
	unit.Get("/:id", handler.Unit.FindByID)
	unit.Put("/:id", handler.Unit.Update)
	unit.Delete("/:id", handler.Unit.Delete)

	category := app.Group("/category", jwtMiddleware)
	category.Post("/", handler.Category.Create)
	category.Get("/", handler.Category.FindAll)
	category.Get("/:id", handler.Category.FindByID)
	category.Put("/:id", handler.Category.Update)
	category.Delete("/:id", handler.Category.Delete)
}
