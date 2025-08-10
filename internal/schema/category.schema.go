package schema

import "github.com/bagusyanuar/app-inventory-be/pkg/pagination"

type CategorySchema struct {
	Name string `json:"name" validate:"required"`
}

type CategoryQuery struct {
	Param string `json:"param" query:"param"`
	pagination.QueryPagination
	pagination.QuerySort
}
