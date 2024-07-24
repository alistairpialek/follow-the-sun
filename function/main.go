package main

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() {
	timezones := [...]string{
		"America/Los_Angeles",
		"Australia/Sydney",
	}

	log.Println("Version: 0.0.1")

	for _, tz := range timezones {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			logError(err)
		}

		timeNow := time.Now().In(loc)

		log.Printf("Time in %s is: %s", tz, timeNow)

		if timeNow.Hour() > 10 && timeNow.Hour() < 16 {
			log.Printf("Time to operate in: %s", tz)
			return
		}
	}

	log.Printf("Looks like a bad time in every region.")
}

func main() {
	if os.Getenv("DEVHOME") != "" {
		HandleRequest()
	} else {
		lambda.Start(HandleRequest)
	}
}

func logError(err error) {
	log.Print(err)
	os.Exit(1)
}
