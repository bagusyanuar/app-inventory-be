package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchAddress struct {
	ID        uuid.UUID
	BranchID  *uuid.UUID
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *BranchAddress) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *BranchAddress) TableName() string {
	return "branch_addresses"
}
