package metrics

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/judegiordano/gogetem/pkg/fibererrors"
)

func Router(r fiber.Router) {
	handler := r.Group("/metrics")
	// middleware
	handler.Use(func(c *fiber.Ctx) error {
		auth := c.Get("x-metrics-auth")
		if auth == "" {
			return fibererrors.Unauthorized(c, errors.New("access not permitted"))
		}
		return c.Next()
	})
	// routes
	handler.Get("/", monitor.New(monitor.Config{Title: "API Monitor"}))
}
