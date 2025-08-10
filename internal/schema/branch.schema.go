package schema

import "github.com/bagusyanuar/app-inventory-be/pkg/pagination"

type BranchSchema struct {
	Name     string                 `json:"name" validate:"required"`
	Address  string                 `json:"address" validate:"required"`
	Contacts []branchContactsSchema `json:"contacts" validate:"required,min=1"`
}

type branchContactsSchema struct {
	Type  string  `json:"type" validate:"required"`
	Name  *string `json:"name"`
	Value string  `json:"value" validate:"required,e164"`
}

type BranchQuery struct {
	Param string `json:"param" query:"param"`
	pagination.QueryPagination
	pagination.QuerySort
}
