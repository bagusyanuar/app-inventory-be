package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Unit struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (e *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *Unit) TableName() string {
	return "units"
}
