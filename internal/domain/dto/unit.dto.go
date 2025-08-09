package dto

import "github.com/bagusyanuar/app-inventory-be/internal/domain/entity"

type UnitDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToUnit(data *entity.Unit) *UnitDTO {
	return &UnitDTO{
		ID:   data.ID.String(),
		Name: data.Name,
	}
}

func ToUnits(data []entity.Unit) []UnitDTO {
	units := make([]UnitDTO, 0)
	for _, datum := range data {
		unit := *ToUnit(&datum)
		units = append(units, unit)
	}
	return units
}
