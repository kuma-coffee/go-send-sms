package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	msg = fmt.Sprintf("Hello, %s!\nNice to meet you!\nHow are you?", "Kuma")
)

func sendMessage(msg, fromPhone, toPhone string, client *twilio.RestClient) {
	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Message SID: %s\n", *response.Sid)
	log.Print("Message sending successful")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		accountSid = os.Getenv("ACCOUNT_SID")
		authToken  = os.Getenv("AUTH_TOKEN")
		fromPhone  = os.Getenv("TWILIO_PHONE_NUMBER")
		toPhone    = os.Getenv("TO_PHONE")
	)

	// init twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	sendMessage(msg, fromPhone, toPhone, client)
}
