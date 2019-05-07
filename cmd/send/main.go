package main

import (
	"flag"
	"log"

	notify "github.com/govau/notify-client-go"
	"github.com/govau/notify-client-go/notifyapi"
)

func check(msg string, err error) {
	if err == nil {
		return
	}

	if nerr, ok := err.(*notifyapi.Error); ok {
		log.Fatalf("%s: Notify API error %d: %T: %v", msg, nerr.Code, nerr, nerr)
		return
	}

	log.Fatalf("%s: %T: %v", msg, err, err)
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
	check("Could not create client", err)

	_, err = client.SendSMS(smsTemplateID, mobile)
	check("Could not send SMS", err)
}
