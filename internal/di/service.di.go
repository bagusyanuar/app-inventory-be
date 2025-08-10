package di

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/service"
)

type ServiceDI struct {
	Auth     service.AuthService
	Unit     service.UnitService
	Category service.CategoryService
}

func MakeDIService(cfg *config.AppConfig, repositoryDI *RepositoryDI) *ServiceDI {
	return &ServiceDI{
		Auth:     service.NewAuthService(repositoryDI.User, cfg),
		Unit:     service.NewUnitService(repositoryDI.Unit, cfg),
		Category: service.NewCategoryService(repositoryDI.Category, cfg),
	}
}
