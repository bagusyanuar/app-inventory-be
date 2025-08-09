package pagination

import "gorm.io/gorm"

const MaxPageSize = 100

type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

func Paginate(db *gorm.DB, page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		switch {
		case pageSize > MaxPageSize:
			pageSize = MaxPageSize
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetTotalPages(totalRows int64, pageSize int) int {
	return int((totalRows + int64(pageSize) - 1) / int64(pageSize))
}

func MakePagination(page, pageSize int, totalRows int64) PaginationMeta {
	return PaginationMeta{
		Page:       page,
		PageSize:   pageSize,
		TotalRows:  totalRows,
		TotalPages: GetTotalPages(totalRows, pageSize),
	}
}
