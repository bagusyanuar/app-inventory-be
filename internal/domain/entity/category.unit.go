package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (e *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *Category) TableName() string {
	return "categories"
}
