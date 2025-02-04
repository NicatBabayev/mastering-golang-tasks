package handlers

import (
	"encoding/json"
	"net/http"
	"session-22/config"
	"session-22/jwt"
	"session-22/model"
)

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (lh *LoginHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		lh.handleLogin(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (lh *LoginHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest model.LoginRequest
	creds, _ := config.ParseCreds()

	userName := creds["CREDS_USERNAME"]
	password := creds["CREDS_PASSWORD"]
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Couldn't decode login request", http.StatusExpectationFailed)
	}
	if loginRequest.UserName != userName && loginRequest.Password != password {
		http.Error(w, "Authorization Failed", http.StatusUnauthorized)
		return
	} else {
		token, err := jwt.CreateToken(userName)
		if err != nil {
			http.Error(w, "Token generation failed", http.StatusInternalServerError)
			return
		}
		loginResp := model.LoginResponse{
			Token: token,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(loginResp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
