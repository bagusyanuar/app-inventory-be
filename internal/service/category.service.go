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
	CategoryService interface {
		FindAll(ctx context.Context, queryParams *schema.CategoryQuery) (*[]dto.CategoryDTO, *pagination.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*dto.CategoryDTO, error)
		Create(ctx context.Context, schema *schema.CategorySchema) (*dto.CategoryDTO, error)
		Update(ctx context.Context, id string, schema *schema.CategorySchema) (*dto.CategoryDTO, error)
		Delete(ctx context.Context, id string) error
	}

	categoryServiceImpl struct {
		CategoryRepository repository.CategoryRepository
		Config             *config.AppConfig
	}
)

func NewCategoryService(
	categoryRepository repository.CategoryRepository,
	config *config.AppConfig,
) CategoryService {
	return &categoryServiceImpl{
		CategoryRepository: categoryRepository,
		Config:             config,
	}
}

// Create implements CategoryService.
func (c *categoryServiceImpl) Create(ctx context.Context, schema *schema.CategorySchema) (*dto.CategoryDTO, error) {
	data := entity.Category{
		Name: schema.Name,
	}

	category, err := c.CategoryRepository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	categoryDTO := dto.ToCategory(category)
	return categoryDTO, nil
}

// FindAll implements CategoryService.
func (c *categoryServiceImpl) FindAll(ctx context.Context, queryParams *schema.CategoryQuery) (*[]dto.CategoryDTO, *pagination.PaginationMeta, error) {
	categories, pagination, err := c.CategoryRepository.FindAll(ctx, queryParams)
	if err != nil {
		return &[]dto.CategoryDTO{}, nil, err
	}
	data := dto.ToCategories(categories)
	return &data, pagination, nil
}

// FindByID implements CategoryService.
func (c *categoryServiceImpl) FindByID(ctx context.Context, id string) (*dto.CategoryDTO, error) {
	category, err := c.CategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	data := dto.ToCategory(category)
	return data, nil
}

// Update implements CategoryService.
func (c *categoryServiceImpl) Update(ctx context.Context, id string, schema *schema.CategorySchema) (*dto.CategoryDTO, error) {
	entry := map[string]any{
		"name": schema.Name,
	}

	category, err := c.CategoryRepository.Update(ctx, id, entry)
	if err != nil {
		return nil, err
	}
	data := dto.ToCategory(category)
	return data, nil
}

// Delete implements CategoryService.
func (c *categoryServiceImpl) Delete(ctx context.Context, id string) error {
	err := c.CategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
