package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/judegiordano/sst_template/api/dev"
	"github.com/judegiordano/sst_template/pkg/helpers"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
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
	fiberLambda = fiberadapter.New(app)
	// app.Listen(":3000")
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
