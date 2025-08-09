package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/app-inventory-be/internal/domain/entity"
	"github.com/bagusyanuar/app-inventory-be/internal/schema"
	"github.com/bagusyanuar/app-inventory-be/pkg/pagination"
	"gorm.io/gorm"
)

type (
	UnitRepository interface {
		FindAll(ctx context.Context, queryParams *schema.UnitQuery) ([]entity.Unit, *pagination.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Unit, error)
		Create(ctx context.Context, unit *entity.Unit) (*entity.Unit, error)
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
	panic("unimplemented")
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
