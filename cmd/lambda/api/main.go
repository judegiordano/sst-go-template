package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/judegiordano/sst_template/internal"
)

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return internal.Lambda.ProxyWithContextV2(ctx, req)
}

func main() {
	if internal.Env.Stage == internal.LocalStage {
		internal.App.Listen(":3000")
		return
	}
	lambda.Start(Handler)
}
