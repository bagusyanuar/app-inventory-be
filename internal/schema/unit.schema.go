package schema

import "github.com/bagusyanuar/app-inventory-be/pkg/pagination"

type UnitSchema struct {
	Name string `json:"name" validate:"required,alphanum"`
}

type UnitQuery struct {
	Param string `json:"param" query:"param"`
	pagination.QueryPagination
	pagination.QuerySort
}
