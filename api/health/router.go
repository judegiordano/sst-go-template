package health

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/judegiordano/gogetem/pkg/fibererrors"
	"github.com/judegiordano/gogetem/pkg/mongo"
	"github.com/judegiordano/sst_template/internal/models"
	"github.com/judegiordano/sst_template/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(c *fiber.Ctx) error {
	health := models.Health{
		Message: "☕",
	}
	inserted, err := health.Save()
	if err != nil {
		return fibererrors.BadRequest(c, err)
	}
	return c.Status(201).JSON(inserted)
}

func list(c *fiber.Ctx) error {
	var limit int64 = 10
	opts := options.FindOptions{
		Limit: &limit,
		Sort:  mongo.Bson{"created_at": -1},
	}
	docs, err := mongo.List[models.Health](mongo.Bson{}, &opts)
	if err != nil {
		return fibererrors.BadRequest(c, err)
	}
	middleware.Cache(c, time.Minute*5)
	return c.JSON(docs)
}

func read(c *fiber.Ctx) error {
	id := c.Params("id")
	doc, err := mongo.ReadById[models.Health](id)
	if err != nil {
		return fibererrors.NotFound(c, err)
	}
	return c.JSON(doc)
}

func delete(c *fiber.Ctx) error {
	id := c.Params("id")
	doc, err := mongo.Delete[models.Health](mongo.Bson{"_id": id})
	if err != nil {
		return fibererrors.NotFound(c, err)
	}
	return c.JSON(doc)
}

func Router(r fiber.Router) {
	handler := r.Group("/health")
	// routes
	handler.Post("/", create)
	handler.Get("/", list)
	handler.Get("/:id", read)
	handler.Delete("/:id", delete)
}
