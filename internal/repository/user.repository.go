package repository

import (
	"context"

	"github.com/bagusyanuar/app-inventory-be/internal/domain/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindByEmail(ctx context.Context, email string) (*entity.User, error)
	}

	userRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

// FindByEmail implements UserRepository.
func (u *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var data *entity.User
	tx := u.DB.WithContext(ctx)
	if err := tx.Where("email = ?", email).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
