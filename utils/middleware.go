package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			json.NewEncoder(w).Encode(Response{false, "Not logged in"})
		} else {
			log.Println(username, password)
			next.ServeHTTP(w, r)
		}
	})
}
