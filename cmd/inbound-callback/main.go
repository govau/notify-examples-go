package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// You should check this matches the one you provided to Notify.
		// If it does not match, then forbid and abort this request.
		log.Printf("Authorization header: %v", r.Header.Get("authorization"))

		if ct := r.Header.Get("content-type"); ct != "application/json" {
			http.Error(
				w,
				fmt.Errorf("got unexpected Content-Type %q, want %q", ct, "application/json").Error(),
				http.StatusBadRequest,
			)
			return
		}

		var req struct {
			ID                string    `json:"id"`
			DateReceived      time.Time `json:"date_received"`
			DestinationNumber string    `json:"destination_number"`
			SourceNumber      string    `json:"source_number"`
			Message           string    `json:"message"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Received inbound message:")
		log.Printf("- ID: %s", req.ID)
		log.Printf("- Date received: %s", req.DateReceived.Format(time.RFC1123))
		log.Printf("- Destination number: %s", req.DestinationNumber)
		log.Printf("- Source number: %s", req.SourceNumber)
		log.Printf("- Message: %s", req.Message)

		w.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":4012", nil)
}
