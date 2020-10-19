package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/dnielsen/campsite/services/api/service"
	"log"
	"net/http"
)

// `/events` POST route. It communicates with the event service only.
func CreateEvent(api service.EventAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i model.EventInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal event input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Request the event service to create the event.
		event, err := api.CreateEvent(i)
		if err != nil {
			log.Printf("Failed to create event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Marshal the event.
		b, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond JSON with the created event.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	}
}
