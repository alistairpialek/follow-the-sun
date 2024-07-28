package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	controlTz   string = "Australia/Sydney"
	rightTz     string = "Australia/Sydney"
	rightRegion string = "ap-southeast-2"
	leftTz      string = "America/Los_Angeles"
	leftRegion  string = "us-west-1"
)

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func HandleRequest(request events.LambdaFunctionURLRequest) (Response, error) {
	log.Println("Version: 0.0.4")
	log.Printf("Request: %+v", request)

	// Control location is used to determine a time of day to run operations.
	controlLoc, err := time.LoadLocation(controlTz)
	if err != nil {
		logError(err)
	}
	controlTimeNow := time.Now().In(controlLoc)

	// Based on the AWS region the lambda runs in, determine the timezone.
	var currentTimezone string
	if os.Getenv("AWS_REGION") == rightRegion {
		currentTimezone = rightTz
	} else if os.Getenv("AWS_REGION") == leftRegion {
		currentTimezone = leftTz
	}

	currentLoc, err := time.LoadLocation(currentTimezone)
	if err != nil {
		logError(err)
	}
	currentTime := time.Now().In(currentLoc)

	// Run the workload during daylight hours.
	if controlTimeNow.Hour() > 10 && controlTimeNow.Hour() < 16 {
		if os.Getenv("AWS_REGION") != rightRegion {
			return Response{
				StatusCode: 301,
				Headers:    map[string]string{"Location": "https://lwptgi26axww5tkkaiso5xknai0erlzm.lambda-url.ap-southeast-2.on.aws"},
			}, nil
		}
	} else {
		if os.Getenv("AWS_REGION") != leftRegion {
			return Response{
				StatusCode: 301,
				Headers:    map[string]string{"Location": "https://fspmcfrkcd7he67r4ns5ueggza0lkjlj.lambda-url.us-west-1.on.aws"},
			}, nil
		}
	}

	// We didn't need to perform a redirect to display where we are.
	return Response{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       fmt.Sprintf("Running in %s as the time now is %s", os.Getenv("AWS_REGION"), currentTime.Format(time.Kitchen)),
	}, nil
}

func main() {
	if os.Getenv("DEVHOME") != "" {
		e := events.LambdaFunctionURLRequest{
			RawPath: "https://url.ap-southeast-2.dev",
		}

		HandleRequest(e)
	} else {
		lambda.Start(HandleRequest)
	}
}

func logError(err error) {
	log.Print(err)
	os.Exit(1)
}
