package handler

import (
	"errors"

	"github.com/bagusyanuar/app-inventory-be/internal/config"
	"github.com/bagusyanuar/app-inventory-be/internal/schema"
	"github.com/bagusyanuar/app-inventory-be/internal/service"
	"github.com/bagusyanuar/app-inventory-be/pkg/exception"
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

func (c *UnitHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(schema.UnitQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeAPIError(ctx, fiber.StatusBadRequest, err)
	}

	messages, err := util.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIErrorValidation(ctx, messages)
	}

	data, pagination, err := c.UnitService.FindAll(ctx.UserContext(), queryParams)
	if err != nil {
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully get units",
		Code:    fiber.StatusOK,
		Data:    data,
		Meta:    pagination,
	})
}

func (c *UnitHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := c.UnitService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return response.MakeAPIError(ctx, fiber.StatusNotFound, err)
		}
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully get unit",
		Code:    fiber.StatusOK,
		Data:    data,
	})
}

func (c *UnitHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := new(schema.UnitSchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeAPIError(ctx, fiber.StatusBadRequest, err)
	}

	messages, err := util.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIErrorValidation(ctx, messages)
	}

	_, err = c.UnitService.Update(ctx.UserContext(), id, request)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return response.MakeAPIError(ctx, fiber.StatusNotFound, err)
		}
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully update unit",
		Code:    fiber.StatusOK,
	})
}

func (c *UnitHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.UnitService.Delete(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return response.MakeAPIError(ctx, fiber.StatusNotFound, err)
		}
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully delete unit",
		Code:    fiber.StatusOK,
	})
}
