package di

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/repository"
)

type RepositoryDI struct {
	User repository.UserRepository
	Unit repository.UnitRepository
}

func MakeDIRepository(cfg *config.AppConfig) *RepositoryDI {
	return &RepositoryDI{
		User: repository.NewUserRepository(cfg.DB),
		Unit: repository.NewUnitRepository(cfg.DB),
	}
}
