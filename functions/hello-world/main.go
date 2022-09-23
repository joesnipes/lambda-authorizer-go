package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.Path == "/hello" {
		return events.APIGatewayProxyResponse{
			Body:       "Hello, World!",
			StatusCode: 200,
		}, nil
	} else if request.Path == "/" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 403,
		}, nil
	}

}

func main() {
	lambda.Start(handler)
}
