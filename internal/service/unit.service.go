package service

import (
	"context"

	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/domain/dto"
	"github.com/bagusyanuar/app-inventory-be/internal/domain/entity"
	"github.com/bagusyanuar/app-inventory-be/internal/repository"
	"github.com/bagusyanuar/app-inventory-be/internal/schema"
	"github.com/bagusyanuar/app-inventory-be/pkg/pagination"
)

type (
	UnitService interface {
		FindAll(ctx context.Context, queryParams *schema.UnitQuery) (*[]dto.UnitDTO, *pagination.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*dto.UnitDTO, error)
		Create(ctx context.Context, schema *schema.UnitSchema) (*dto.UnitDTO, error)
	}

	unitServiceImpl struct {
		UnitRepository repository.UnitRepository
		Config         *config.AppConfig
	}
)

func NewUnitService(
	unitRepository repository.UnitRepository,
	config *config.AppConfig,
) UnitService {
	return &unitServiceImpl{
		UnitRepository: unitRepository,
		Config:         config,
	}
}

// Create implements UnitService.
func (u *unitServiceImpl) Create(ctx context.Context, schema *schema.UnitSchema) (*dto.UnitDTO, error) {
	data := entity.Unit{
		Name: schema.Name,
	}

	unit, err := u.UnitRepository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	unitDTO := dto.ToUnit(unit)
	return unitDTO, nil
}

// FindAll implements UnitService.
func (u *unitServiceImpl) FindAll(ctx context.Context, queryParams *schema.UnitQuery) (*[]dto.UnitDTO, *pagination.PaginationMeta, error) {
	panic("unimplemented")
}

// FindByID implements UnitService.
func (u *unitServiceImpl) FindByID(ctx context.Context, id string) (*dto.UnitDTO, error) {
	panic("unimplemented")
}
