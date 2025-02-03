package main

import (
	"github.com/judegiordano/sst_template/TEST/internal"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	handler := internal.Server()

	lambda.Start(httpadapter.NewV2(handler).ProxyWithContext)

	// lambda.Start(httpadapter.New(handler).ProxyWithContext)
}
