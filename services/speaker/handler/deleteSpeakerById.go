package handler

import (
	"github.com/dnielsen/campsite/services/speaker/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// `/` DELETE route.
func DeleteSpeakerById(datastore service.SpeakerAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Delete the speaker from the database.
		if err := datastore.DeleteSpeakerById(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.NotFound(w, r)
			return
		}
		// Respond that the speaker deletion has been successful.
		w.WriteHeader(http.StatusNoContent)
	}
}
