package pagination

import (
	"fmt"

	"gorm.io/gorm"
)

func SortScope(sort, order string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		sort := fmt.Sprintf("%s %s", sort, order)
		return tx.Order(sort)
	}
}

func GetSortField(sortKey, defaultField string, fieldMap map[string]string) string {
	if field, ok := fieldMap[sortKey]; ok {
		return field
	}
	return defaultField
}

func GetOrder(order string) string {
	val := "ASC"
	if order == "DESC" {
		val = order
	}
	return val
}
