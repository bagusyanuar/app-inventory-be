package di

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/repository"
)

type RepositoryDI struct {
	Unit repository.UnitRepository
}

func MakeDIRepository(cfg *config.AppConfig) *RepositoryDI {
	return &RepositoryDI{
		Unit: repository.NewUnitRepository(cfg.DB),
	}
}
