package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HelixY2J/echo/common/db"
	"github.com/HelixY2J/echo/common/models"
	"github.com/HelixY2J/echo/common/publisher"
)

func PublishHandler(pub *publisher.Publisher, store *db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		//	var notification map[string]interface{}
		var notification models.Notification
		if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err := store.InsertNotification(notification)
		if err != nil {
			log.Printf("failed to save notif: %v", err)
			http.Error(w, "Failed to save notification", http.StatusInternalServerError)
			return
		}

		notificationBytes, _ := json.Marshal(notification)
		//log.Printf("Publishing to nat: %s", string(notificationBytes))

		if err := pub.PublishNotification(notificationBytes); err != nil {
			http.Error(w, "Oops failed to PuBlIsh Notification", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}
