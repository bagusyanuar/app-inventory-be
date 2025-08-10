package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/bagusyanuar/app-inventory-be/internal/domain/entity"
	"github.com/bagusyanuar/app-inventory-be/internal/schema"
	"github.com/bagusyanuar/app-inventory-be/pkg/exception"
	"github.com/bagusyanuar/app-inventory-be/pkg/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	UnitRepository interface {
		FindAll(ctx context.Context, queryParams *schema.UnitQuery) ([]entity.Unit, *pagination.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Unit, error)
		Create(ctx context.Context, unit *entity.Unit) (*entity.Unit, error)
		Update(ctx context.Context, id string, entry map[string]any) (*entity.Unit, error)
		Delete(ctx context.Context, id string) error
	}

	unitRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepositoryImpl{
		DB: db,
	}
}

// Create implements UnitRepository.
func (u *unitRepositoryImpl) Create(ctx context.Context, unit *entity.Unit) (*entity.Unit, error) {
	tx := u.DB.WithContext(ctx)
	if err := tx.Create(&unit).Error; err != nil {
		return nil, err
	}
	return unit, nil
}

// FindAll implements UnitRepository.
func (u *unitRepositoryImpl) FindAll(ctx context.Context, queryParams *schema.UnitQuery) ([]entity.Unit, *pagination.PaginationMeta, error) {
	tx := u.DB.WithContext(ctx)
	baseQuery := u.baseQuery(tx, queryParams)

	var totalRows int64
	if err := baseQuery.
		Model(&entity.Unit{}).
		Count(&totalRows).Error; err != nil {
		return []entity.Unit{}, nil, err
	}

	var data []entity.Unit

	sortFieldMap := map[string]string{
		"name": "name",
	}
	sort := pagination.GetSortField(queryParams.Sort, "name", sortFieldMap)
	order := pagination.GetOrder(queryParams.Order)
	if err := baseQuery.
		Scopes(
			pagination.SortScope(sort, order),
			pagination.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).Error; err != nil {
		return []entity.Unit{}, nil, err
	}

	pagination := pagination.MakePagination(queryParams.Page, queryParams.PageSize, totalRows)
	return data, &pagination, nil
}

// FindByID implements UnitRepository.
func (u *unitRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Unit, error) {
	var data *entity.Unit
	tx := u.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}
	return data, nil
}

// Update implements UnitRepository.
func (u *unitRepositoryImpl) Update(ctx context.Context, id string, entry map[string]any) (*entity.Unit, error) {
	var data *entity.Unit
	tx := u.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	if err := tx.Model(&data).
		Omit(clause.Associations).Where("id = ?", id).
		Updates(&entry).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Delete implements UnitRepository.
func (u *unitRepositoryImpl) Delete(ctx context.Context, id string) error {
	var data *entity.Unit
	tx := u.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrRecordNotFound
		}
		return err
	}

	if err := tx.Omit(clause.Associations).
		Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (u *unitRepositoryImpl) baseQuery(tx *gorm.DB, queryParams *schema.UnitQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)

	return tx.
		Scopes(
			u.filterByParam(param),
		)
}

func (u *unitRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}
		return tx.
			Where("name ILIKE ?", param)
	}
}
