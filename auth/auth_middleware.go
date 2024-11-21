package auth

import (
	"log"
	"net/http"

	pkgResource "github.com/DestroyerAlpha/COMSW4156-AuthDemo/pkg/resource"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthMiddleware")
		log.Println(r.Header)
		if isSkipMiddleware(r.Header.Get("resource")) {
			log.Println("skip auth middleware")
			h.ServeHTTP(w, r)
			return
		}
		accessToken := r.Header.Get("access_token")
		if accessToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var claims Claims
		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("SuperSecretKey"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
		}
		r.Header.Set("user_id", claims.UserId)
		h.ServeHTTP(w, r)
	})
}

func isSkipMiddleware(resource string) bool {
	if resource == pkgResource.CreateUser || resource == pkgResource.Authenticate {
		return true
	}
	return false
}
