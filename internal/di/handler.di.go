package di

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/http/handler"
)

type HandlerDI struct {
	Auth *handler.AuthHandler
	Unit *handler.UnitHandler
}

func MakeDIHandler(cfg *config.AppConfig, serviceDI *ServiceDI) *HandlerDI {
	return &HandlerDI{
		Auth: handler.NewAuthHandler(serviceDI.Auth, cfg),
		Unit: handler.NewUnitHandler(serviceDI.Unit, cfg),
	}
}
