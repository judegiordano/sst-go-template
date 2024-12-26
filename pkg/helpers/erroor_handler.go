package helpers

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
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
	log.Error(code, message)
	return ctx.Status(code).JSON(ErrorResponse{
		Error: message,
	})
}
