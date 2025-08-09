package config

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func fiberConfig(viper *viper.Viper) fiber.Config {
	config := fiber.Config{
		TrustedProxies:        []string{"0.0.0.0/0"},
		DisableStartupMessage: viper.GetString("APP_ENV") == "production",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			switch code {
			case fiber.StatusBadRequest,
				fiber.StatusUnauthorized,
				fiber.StatusForbidden,
				fiber.StatusUnprocessableEntity:
				return c.Status(code).JSON(fiber.Map{
					"code":    code,
					"message": err.Error(),
				})
			case fiber.StatusNotFound:
				return c.Status(code).JSON(fiber.Map{
					"code":    code,
					"message": "route not found",
				})
			default:
				return c.Status(code).JSON(fiber.Map{
					"code":    code,
					"message": "internal server error",
				})
			}
		},
	}
	return config
}

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiberConfig(viper))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	return app
}
