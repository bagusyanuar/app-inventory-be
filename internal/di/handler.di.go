package di

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/http/handler"
)

type HandlerDI struct {
	Unit *handler.UnitHandler
}

func MakeDIHandler(cfg *config.AppConfig, serviceDI *ServiceDI) *HandlerDI {
	return &HandlerDI{
		Unit: handler.NewUnitHandler(serviceDI.Unit, cfg),
	}
}
