package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/pkg/constant"
	"github.com/bagusyanuar/app-inventory-be/pkg/exception"
	"github.com/bagusyanuar/app-inventory-be/pkg/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func VerifyJWT(cfg *config.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if errors.Is(err, jwtware.ErrJWTMissingOrMalformed) {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: exception.ErrMissingOrMalformed.Error(),
				})
			}

			if strings.Contains(err.Error(), "token is expired") {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: exception.ErrTokenExpired.Error(),
				})
			}

			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Code:    fiber.StatusUnauthorized,
				Message: exception.ErrUnauthorized.Error(),
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			sub, ok := claims["sub"].(string)
			if !ok || sub == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: exception.ErrClaimToken.Error(),
				})
			}

			userID, err := uuid.Parse(sub)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
					Code:    fiber.StatusUnauthorized,
					Message: exception.ErrInvalidSubjectFormat.Error(),
				})
			}

			ctx := context.WithValue(c.UserContext(), constant.UserIDKey, userID)
			c.SetUserContext(ctx)
			return c.Next()
		},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(cfg.JWT.Secret),
		},
	})
}
