package handler

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/domain/dto"
	"github.com/bagusyanuar/app-inventory-be/internal/schema"
	"github.com/bagusyanuar/app-inventory-be/internal/service"
	"github.com/bagusyanuar/app-inventory-be/pkg/response"
	"github.com/bagusyanuar/app-inventory-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService service.AuthService
	Config      *config.AppConfig
}

func NewAuthHandler(
	authService service.AuthService,
	config *config.AppConfig,
) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Config:      config,
	}
}

func (c *AuthHandler) Login(ctx *fiber.Ctx) error {
	request := new(schema.LoginSchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		})
	}

	messages, err := util.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: fiber.ErrUnprocessableEntity.Error(),
			Code:    fiber.StatusUnprocessableEntity,
			Data:    messages,
		})
	}

	data, err := c.AuthService.Login(ctx.UserContext(), request)
	if err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: err.Error(),
			Code:    fiber.StatusInternalServerError,
		})
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[*dto.LoginDTO]{
		Message: "successfully login",
		Code:    fiber.StatusOK,
		Data:    data,
	})
}
