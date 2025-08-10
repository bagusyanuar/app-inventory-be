package response

import (
	"github.com/bagusyanuar/app-inventory-be/pkg/exception"
	"github.com/gofiber/fiber/v2"
)

type (
	APIResponse[T any] struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    T      `json:"data,omitempty"`
		Meta    any    `json:"meta,omitempty"`
	}

	APIResponseOptions[T any] struct {
		Message string
		Data    T
		Meta    any
	}
)

func MakeAPIResponse[T any](ctx *fiber.Ctx, res APIResponse[T]) error {
	return ctx.Status(res.Code).JSON(res)
}

func MakeAPIError(ctx *fiber.Ctx, code int, err error) error {
	res := APIResponse[any]{
		Code:    code,
		Message: err.Error(),
	}
	return ctx.Status(code).JSON(res)
}

func MakeAPIErrorValidation(ctx *fiber.Ctx, messages map[string][]string) error {
	code := fiber.StatusUnprocessableEntity
	res := APIResponse[map[string][]string]{
		Code:    code,
		Message: exception.ErrUnprocessableEntity.Error(),
		Data:    messages,
	}
	return ctx.Status(code).JSON(res)
}
