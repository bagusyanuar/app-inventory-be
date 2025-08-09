package pagination

type QueryPagination struct {
	Page     int `json:"page" query:"page" validate:"required"`
	PageSize int ` json:"page_size" query:"page_size" validate:"required"`
}

type QuerySort struct {
	Sort  string `json:"sort" query:"sort"`
	Order string ` json:"order" query:"order"`
}
