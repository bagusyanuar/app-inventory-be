package response

import "github.com/gofiber/fiber/v2"

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
