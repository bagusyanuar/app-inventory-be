package bootstrap

import (
	"fmt"

	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/di"
	"github.com/bagusyanuar/app-inventory-be/internal/http"
)

func initialize() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	logger := config.NewLogger(viper)
	defer logger.Sync()
	cfgDB := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(cfgDB)
	cfgJWT := config.NewJWTManager(viper)
	validator := config.NewValidator()

	return &config.AppConfig{
		App:       app,
		Viper:     viper,
		DB:        db,
		Logger:    logger,
		JWT:       cfgJWT,
		Validator: validator,
	}
}

func Start() {
	cfg := initialize()

	// initialize dependency injection
	repositoryDI := di.MakeDIRepository(cfg)
	serviceDI := di.MakeDIService(cfg, repositoryDI)
	handlerDI := di.MakeDIHandler(cfg, serviceDI)
	http.NewRouter(cfg, handlerDI)
	envPort := cfg.Viper.GetString("APP_PORT")
	port := fmt.Sprintf(":%s", envPort)
	server := cfg.App
	fmt.Println("Fiber server running on", port)
	if err := server.Listen(port); err != nil {
		panic(err)
	}
}
