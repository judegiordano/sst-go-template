package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	adaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"

	"github.com/judegiordano/sst_template/internal"
)

func main() {
	app := internal.Server()
	lam := adaptor.New(app)
	lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		return lam.ProxyWithContextV2(ctx, req)
	})
}
