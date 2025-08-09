package main

import (
	"github.com/bagusyanuar/app-inventory-be/database/seed"
	"github.com/bagusyanuar/app-inventory-be/internal/config"
)

func main() {
	viper := config.NewViper()
	dbConfig := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(dbConfig)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	seed.Seed(db)
}
