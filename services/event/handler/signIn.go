package handler

import (
	"campsite/pkg/model"
	"campsite/services/event/service"
	"encoding/json"
	"log"
	"net/http"
)

func SignIn(api service.AuthAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the body.
		var i model.SignInInput
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			log.Printf("Failed to unmarshal sign in input")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate the credentials match the user.
		token, err := api.SignIn(i)
		if err != nil {
			log.Printf("Failed to validate user: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Respond plain text with the token. We might change the response later,
		// to some json object.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	}
}
