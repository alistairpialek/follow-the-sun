package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func HandleRequest() (Response, error) {
	return Response{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       fmt.Sprintf("Running in %s", os.Getenv("AWS_REGION")),
	}, nil
}

func main() {
	if os.Getenv("DEVHOME") != "" {
		HandleRequest()
	} else {
		lambda.Start(HandleRequest)
	}
}
