package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchContact struct {
	ID        uuid.UUID
	BranchID  *uuid.UUID
	Type      string
	Name      *string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *BranchContact) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *BranchContact) TableName() string {
	return "branch_contacts"
}
