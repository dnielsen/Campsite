package handler

import (
	"github.com/dnielsen/campsite/services/session/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteSessionById(datastore service.SessionAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the id parameter.
		vars := mux.Vars(r)
		id := vars[ID]
		// Delete the session from the database.
		if err := datastore.DeleteSessionById(id); err != nil {
			log.Printf("Failed to delete session: %v", err)
			http.NotFound(w, r)
			return
		}
		// Respond that the session deletion has been successful.
		w.WriteHeader(http.StatusNoContent)
	}
}
