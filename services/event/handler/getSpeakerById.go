package handler

import (
	"encoding/json"
	"github.com/dnielsen/campsite/services/event/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/speakers/{id}` GET route. It communicates with the speaker service only.
func GetSpeakerById(api service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]

		// Get the speaker from the speaker service.
		speaker, err := api.GetSpeakerById(id)
		if err != nil {
			log.Printf("Failed to get speaker: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Marshal the speaker.
		speakerBytes, err := json.Marshal(speaker)
		if err != nil {
			log.Printf("Failed to marshal speaker: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond JSON with the speaker.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(speakerBytes)
	}
}
