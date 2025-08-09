package di

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/service"
)

type ServiceDI struct {
	Unit service.UnitService
}

func MakeDIService(cfg *config.AppConfig, repositoryDI *RepositoryDI) *ServiceDI {
	return &ServiceDI{
		Unit: service.NewUnitService(repositoryDI.Unit, cfg),
	}
}
