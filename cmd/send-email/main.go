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
	var templateID string
	var email string
	var replyTo string

	flag.StringVar(&apiKey, "api-key", "", "Notify API key")
	flag.StringVar(&templateID, "template-id", "", "Notify email template ID, e.g. c1ad8967-41ae-4013-bdc1-af29d2ef3ce9")
	flag.StringVar(&email, "email", "", "Recipient's email address")
	flag.StringVar(&replyTo, "reply-to", "", "Email reply to ID e.g. c1ad8967-41ae-4013-bdc1-af29d2ef3ce9")

	flag.Parse()

	client, err := notify.NewClient(apiKey)
	check("Could not create client", err)

	options := []notify.SendEmailOption{
		notify.Personalisation{
			{"name", "Kim"},
		},
	}

	if replyTo != "" {
		options = append(options, notify.EmailReplyToID(replyTo))
	}

	_, err = client.SendEmail(templateID, email, options...)
	check("Could not send message", err)
}
