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

type CategoryHandler struct {
	CategoryService service.CategoryService
	Config          *config.AppConfig
}

func NewCategoryHandler(
	categoryService service.CategoryService,
	cfg *config.AppConfig,
) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
		Config:          cfg,
	}
}

func (c *CategoryHandler) Create(ctx *fiber.Ctx) error {
	request := new(schema.CategorySchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeAPIError(ctx, fiber.StatusBadRequest, err)
	}

	messages, err := util.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIErrorValidation(ctx, messages)
	}

	_, err = c.CategoryService.Create(ctx.UserContext(), request)
	if err != nil {
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}

	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully create new category",
		Code:    fiber.StatusCreated,
	})
}

func (c *CategoryHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(schema.CategoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeAPIError(ctx, fiber.StatusBadRequest, err)
	}

	messages, err := util.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIErrorValidation(ctx, messages)
	}

	data, pagination, err := c.CategoryService.FindAll(ctx.UserContext(), queryParams)
	if err != nil {
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully get categories",
		Code:    fiber.StatusOK,
		Data:    data,
		Meta:    pagination,
	})
}

func (c *CategoryHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := c.CategoryService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return response.MakeAPIError(ctx, fiber.StatusNotFound, err)
		}
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully get category",
		Code:    fiber.StatusOK,
		Data:    data,
	})
}

func (c *CategoryHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := new(schema.CategorySchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeAPIError(ctx, fiber.StatusBadRequest, err)
	}

	messages, err := util.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIErrorValidation(ctx, messages)
	}

	_, err = c.CategoryService.Update(ctx.UserContext(), id, request)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return response.MakeAPIError(ctx, fiber.StatusNotFound, err)
		}
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully update category",
		Code:    fiber.StatusOK,
	})
}

func (c *CategoryHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.CategoryService.Delete(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return response.MakeAPIError(ctx, fiber.StatusNotFound, err)
		}
		return response.MakeAPIError(ctx, fiber.StatusInternalServerError, err)
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully delete category",
		Code:    fiber.StatusOK,
	})
}
