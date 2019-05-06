package main

import (
	"flag"
	"log"

	notify "github.com/govau/notify-client-go"
)

func check(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %T: %v", msg, err, err)
	}
}

func main() {
	var apiKey string
	var smsTemplateID string
	var mobile string

	flag.StringVar(&apiKey, "api-key", "", "Notify API key")
	flag.StringVar(&smsTemplateID, "sms-template-id", "", "Notify SMS template ID")
	flag.StringVar(&mobile, "mobile", "", "Recipient's mobile number")

	flag.Parse()

	client, err := notify.NewClient(apiKey)
	check("could not create client", err)

	_, err = client.SendSMS(smsTemplateID, mobile)
	check("could not send SMS", err)
}
