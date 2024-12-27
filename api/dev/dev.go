package dev

import (
	"github.com/gofiber/fiber/v2"
)

type Health struct {
	Message string `json:"message"`
}

func healthCheck(c *fiber.Ctx) error {
	c.Status(418)
	return c.JSON(Health{Message: "☕"})
}

func Router(r fiber.Router) {
	handler := r.Group("/dev")
	// routes
	handler.Get("/health", healthCheck)
}
