package dto

import "github.com/bagusyanuar/app-inventory-be/internal/domain/entity"

type CategoryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToCategory(data *entity.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:   data.ID.String(),
		Name: data.Name,
	}
}

func ToCategories(data []entity.Category) []CategoryDTO {
	categories := make([]CategoryDTO, 0)
	for _, datum := range data {
		category := *ToCategory(&datum)
		categories = append(categories, category)
	}
	return categories
}
