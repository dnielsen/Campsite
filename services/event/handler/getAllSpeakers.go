package handler

import (
	"campsite/pkg/model"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllSpeakers(api model.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all speakers from the speaker service.
		speakers, err := api.GetAllSpeakers()
		if err != nil {
			log.Printf("Failed to get speakers: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Marshal the speakers.
		speakerBytes, err := json.Marshal(speakers)
		if err != nil {
			log.Printf("Failed to marshal sessions: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond json with the speakers.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
