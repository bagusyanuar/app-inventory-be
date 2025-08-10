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
)

type (
	BranchRepository interface {
		FindAll(ctx context.Context, queryParams *schema.BranchQuery) ([]entity.Branch, *pagination.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Branch, error)
		Create(ctx context.Context, branch *entity.Branch) (*entity.Branch, error)
		Update(ctx context.Context, id string, entry map[string]any) (*entity.Branch, error)
		Delete(ctx context.Context, id string) error
	}

	branchRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &branchRepositoryImpl{
		DB: db,
	}
}

// Create implements BranchRepository.
func (b *branchRepositoryImpl) Create(ctx context.Context, branch *entity.Branch) (*entity.Branch, error) {
	tx := b.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&branch).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return branch, nil
}

// FindAll implements BranchRepository.
func (b *branchRepositoryImpl) FindAll(ctx context.Context, queryParams *schema.BranchQuery) ([]entity.Branch, *pagination.PaginationMeta, error) {
	tx := b.DB.WithContext(ctx)
	baseQuery := b.baseQuery(tx, queryParams)

	var totalRows int64
	if err := baseQuery.
		Model(&entity.Branch{}).
		Count(&totalRows).Error; err != nil {
		return []entity.Branch{}, nil, err
	}

	var data []entity.Branch

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
		return []entity.Branch{}, nil, err
	}

	pagination := pagination.MakePagination(queryParams.Page, queryParams.PageSize, totalRows)
	return data, &pagination, nil
}

// FindByID implements BranchRepository.
func (b *branchRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Branch, error) {
	var data *entity.Branch
	tx := b.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}
	return data, nil
}

// Update implements BranchRepository.
func (b *branchRepositoryImpl) Update(ctx context.Context, id string, entry map[string]any) (*entity.Branch, error) {
	var data *entity.Branch
	tx := b.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	err := tx.Transaction(func(tx *gorm.DB) error {

		// 1. map for parent table
		mainEntry := make(map[string]any)
		if v, ok := entry["name"]; ok {
			mainEntry["name"] = v
		}

		// 2. update parent table
		if len(mainEntry) > 0 {
			if err := tx.Model(&entity.Branch{}).
				Where("id = ?", id).
				Updates(entry).Error; err != nil {
				return err
			}
		}

		// 3. update address table
		if v, ok := entry["address"]; ok {
			if addressMap, ok := v.(map[string]any); ok {
				if err := tx.Model(&entity.BranchAddress{}).
					Where("branch_id = ?", id).
					Updates(addressMap).Error; err != nil {
					return err
				}
			}
		}

		// 4. update contatcs table

		// if contacts, ok := entry["contacts"]; ok {

		// }
		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements BranchRepository.
func (b *branchRepositoryImpl) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (c *branchRepositoryImpl) baseQuery(tx *gorm.DB, queryParams *schema.BranchQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)

	return tx.
		Preload("Address").
		Preload("Contacts").
		Scopes(
			c.filterByParam(param),
		)
}

func (c *branchRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}
		return tx.
			Where("name ILIKE ?", param)
	}
}
