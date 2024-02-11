package middleware

import (
	"net/http"

	utils "github.com/personal-finance-app/utils/jwt"
)

func Authorization(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not authorised"))
			return
		}

		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorised"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
