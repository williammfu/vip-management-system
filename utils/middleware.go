package utils

import (
	"encoding/json"
	"net/http"

	"github.com/williammfu/vip-management-system/config"
	"gorm.io/gorm"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type Staff struct {
	Name  string
	Token string
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, token, ok := r.BasicAuth()
		if !ok {
			json.NewEncoder(w).Encode(Response{false, "Not logged in"})
		} else {
			var staff Staff
			result := db.Where("name = ?", username).First(&staff)
			if result.Error != nil {
				json.NewEncoder(w).Encode(Response{false, "Error on login"})
			} else {
				if token != staff.Token {
					json.NewEncoder(w).Encode(Response{false, "Wrong Credentials"})
				} else {
					next.ServeHTTP(w, r)
				}
			}
		}
	})
}
