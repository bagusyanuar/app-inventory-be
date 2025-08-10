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
	CategoryRepository interface {
		FindAll(ctx context.Context, queryParams *schema.CategoryQuery) ([]entity.Category, *pagination.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Category, error)
		Create(ctx context.Context, category *entity.Category) (*entity.Category, error)
		Update(ctx context.Context, id string, entry map[string]any) (*entity.Category, error)
		Delete(ctx context.Context, id string) error
	}

	categoryRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		DB: db,
	}
}

// Create implements CategoryRepository.
func (c *categoryRepositoryImpl) Create(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	tx := c.DB.WithContext(ctx)
	if err := tx.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

// FindAll implements CategoryRepository.
func (c *categoryRepositoryImpl) FindAll(ctx context.Context, queryParams *schema.CategoryQuery) ([]entity.Category, *pagination.PaginationMeta, error) {
	tx := c.DB.WithContext(ctx)
	baseQuery := c.baseQuery(tx, queryParams)

	var totalRows int64
	if err := baseQuery.
		Model(&entity.Category{}).
		Count(&totalRows).Error; err != nil {
		return []entity.Category{}, nil, err
	}

	var data []entity.Category

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
		return []entity.Category{}, nil, err
	}

	pagination := pagination.MakePagination(queryParams.Page, queryParams.PageSize, totalRows)
	return data, &pagination, nil
}

// FindByID implements CategoryRepository.
func (c *categoryRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Category, error) {
	var data *entity.Category
	tx := c.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}
	return data, nil
}

// Update implements CategoryRepository.
func (c *categoryRepositoryImpl) Update(ctx context.Context, id string, entry map[string]any) (*entity.Category, error) {
	var data *entity.Category
	tx := c.DB.WithContext(ctx)
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

// Delete implements CategoryRepository.
func (c *categoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	var data *entity.Category
	tx := c.DB.WithContext(ctx)
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

func (c *categoryRepositoryImpl) baseQuery(tx *gorm.DB, queryParams *schema.CategoryQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	return tx.
		Scopes(
			c.filterByParam(param),
		)
}

func (c *categoryRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}
		return tx.
			Where("name ILIKE ?", param)
	}
}
