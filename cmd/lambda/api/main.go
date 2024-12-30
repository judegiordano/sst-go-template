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
	"github.com/judegiordano/sst_template/internal"
	"github.com/judegiordano/sst_template/pkg/errors"
	"github.com/judegiordano/sst_template/pkg/logger"
)

var fiberLambda *fiberadapter.FiberLambda
var app *fiber.App

func init() {
	logger.SetLogLevel(internal.Env.LogLevel)
	app = fiber.New(fiber.Config{
		ErrorHandler:      errors.ErrorHandler,
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
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, req)
}

func main() {
	if internal.Env.Stage == internal.LocalStage {
		app.Listen(":3000")
		return
	}
	lambda.Start(Handler)
}
