package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth/dao"

	"github.com/golang-jwt/jwt/v5"
)

type AuthenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	AccessToken string `json:"access_token"`
}

type AccessToken struct {
	UserId string `json:"user_id"`
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	var req AuthenticateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()
	if !dao.IsValidUsernameAndPassword(req.Username, req.Password) {
		http.Error(w, fmt.Sprintf("Invalid credentials: %v", err), http.StatusUnauthorized)
		return
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserId: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}).SignedString([]byte("SuperSecretKey"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to sign JWT: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(AuthenticateResponse{
		AccessToken: token,
	}); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
