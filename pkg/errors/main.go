package errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/judegiordano/sst_template/pkg/logger"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}
	logger.Error(code, message)
	return ctx.Status(code).JSON(ErrorResponse{
		Error: message,
	})
}
