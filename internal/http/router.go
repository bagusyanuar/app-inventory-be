package http

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/di"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerDI) {
	app := cfg.App

	auth := app.Group("/auth")
	auth.Post("/login", handler.Auth.Login)

	unit := app.Group("/unit")
	unit.Post("/", handler.Unit.Create)
}
