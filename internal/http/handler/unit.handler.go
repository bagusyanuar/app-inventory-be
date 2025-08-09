package handler

import (
	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/schema"
	"github.com/bagusyanuar/app-inventory-be/internal/service"
	"github.com/bagusyanuar/app-inventory-be/pkg/response"
	"github.com/bagusyanuar/app-inventory-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type UnitHandler struct {
	UnitService service.UnitService
	Config      *config.AppConfig
}

func NewUnitHandler(
	unitService service.UnitService,
	cfg *config.AppConfig,
) *UnitHandler {
	return &UnitHandler{
		UnitService: unitService,
		Config:      cfg,
	}
}

func (c *UnitHandler) Create(ctx *fiber.Ctx) error {
	request := new(schema.UnitSchema)
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

	_, err = c.UnitService.Create(ctx.UserContext(), request)
	if err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: err.Error(),
			Code:    fiber.StatusInternalServerError,
		})
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully create new unit",
		Code:    fiber.StatusCreated,
	})
}
