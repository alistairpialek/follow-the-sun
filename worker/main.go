package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (string, error) {
	message := fmt.Sprintf("Running in %s", os.Getenv("AWS_REGION"))
	return message, nil
}

func main() {
	if os.Getenv("DEVHOME") != "" {
		HandleRequest()
	} else {
		lambda.Start(HandleRequest)
	}
}
