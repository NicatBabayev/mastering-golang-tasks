package middleware

import (
	"fmt"
	"net/http"
	"session-22/jwt"
)

const (
	AuthHeader = "Authorization"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" || r.Method == "PUT" || r.Method == "GET" {
			token := r.Header.Get(AuthHeader)
			if token == "" {
				http.Error(w, "Authorization failed", http.StatusUnauthorized)
				return
			}
			_, err := jwt.Verify(token)
			if err == nil {
				next(w, r)
			} else {
				fmt.Println(err)
				http.Error(w, "Authorization failed", http.StatusUnauthorized)
			}
		} else {
			next(w, r)
		}
	}
}
