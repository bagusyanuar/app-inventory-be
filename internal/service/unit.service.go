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
		Update(ctx context.Context, id string, schema *schema.UnitSchema) (*dto.UnitDTO, error)
		Delete(ctx context.Context, id string) error
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
	units, pagination, err := u.UnitRepository.FindAll(ctx, queryParams)
	if err != nil {
		return &[]dto.UnitDTO{}, nil, err
	}
	data := dto.ToUnits(units)
	return &data, pagination, nil
}

// FindByID implements UnitService.
func (u *unitServiceImpl) FindByID(ctx context.Context, id string) (*dto.UnitDTO, error) {
	unit, err := u.UnitRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	data := dto.ToUnit(unit)
	return data, nil
}

// Update implements UnitService.
func (u *unitServiceImpl) Update(ctx context.Context, id string, schema *schema.UnitSchema) (*dto.UnitDTO, error) {
	entry := map[string]any{
		"name": schema.Name,
	}

	unit, err := u.UnitRepository.Update(ctx, id, entry)
	if err != nil {
		return nil, err
	}
	data := dto.ToUnit(unit)
	return data, nil
}

// Delete implements UnitService.
func (u *unitServiceImpl) Delete(ctx context.Context, id string) error {
	err := u.UnitRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
