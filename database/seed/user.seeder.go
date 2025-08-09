package seed

import (
	"errors"
	"log"

	"github.com/bagusyanuar/app-inventory-be/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("users") {
		log.Println("⛔ Table users not found, seeding cancelled.")
		return
	}

	email := "superdev@web.id"
	username := "superdev"
	password := "@Superdev123"

	hash, errHashed := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashed != nil {
		log.Printf("❌ failed to hash password: %v", errHashed)
		return
	}

	data := entity.User{
		Email:    email,
		Username: username,
		Password: string(hash),
	}

	var user *entity.User
	err := db.Where("email = ?", email).First(&user).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		if err := db.Create(&data).Error; err != nil {
			log.Printf("❌ failed insert user: %v", err)
			return
		}
		log.Println("✅ successfully insert user (insert)")
	case err != nil:
		log.Printf("❌ failed to find user: %v", err)
		return
	default:
		if errUpdate := db.Model(&user).Updates(&data).Error; errUpdate != nil {
			log.Printf("❌ failed to update seed user: %v", err)
			return
		}
		log.Println("✅ successfully update user (update)")
	}
}
