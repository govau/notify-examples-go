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
			ID               string    `json:"id"`
			Reference        string    `json:"reference"`
			To               string    `json:"to"`
			Status           string    `json:"status"`
			NotificationType string    `json:"notification_type"`
			CreatedAt        time.Time `json:"created_at"`
			CompletedAt      time.Time `json:"completed_at"`
			SentAt           time.Time `json:"sent_at"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Delivery status update:")
		log.Printf("- ID: %s", req.ID)
		log.Printf("- Reference: %s", req.Reference)
		log.Printf("- To: %s", req.To)
		log.Printf("- Status: %s", req.Status)
		log.Printf("- Notification type: %s", req.NotificationType)
		log.Printf("- Date created: %s", req.CreatedAt.Format(time.RFC1123))
		log.Printf("- Date completed: %s", req.CompletedAt.Format(time.RFC1123))
		log.Printf("- Date sent: %s", req.SentAt.Format(time.RFC1123))

		w.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":4011", nil)
}
