package middleware

import (
	"net/http"
	"strings"

	utils "github.com/personal-finance-app/utils/jwt"
)

func Authorization(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorised"))
			return
		}

		_, err := utils.VerifyToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorised"))
			return
		}

		handler.ServeHTTP(w, r)
	}
}
