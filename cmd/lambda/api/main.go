package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/judegiordano/sst_template/api/dev"
	"github.com/judegiordano/sst_template/pkg/helpers"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler:      helpers.ErrorHandler,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		EnablePrintRoutes: false,
	})
	// routes
	api := app.Group("/api")
	dev.Router(api)
	// middleware
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(cors.New())
	lambda.Start(app)
	// app.Listen(":3000")
}
