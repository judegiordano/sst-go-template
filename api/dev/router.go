package dev

import (
	"github.com/gofiber/fiber/v2"
)

type Ping struct {
	Ok bool `json:"ok"`
}

func ping(c *fiber.Ctx) error {
	return c.JSON(Ping{Ok: true})
}

func Router(r fiber.Router) {
	handler := r.Group("/dev")
	// routes
	handler.Get("/ping", ping)
}
