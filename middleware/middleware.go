package middleware

import (
	"net/http"

	utils "github.com/personal-finance-app/utils/jwt"
)

func Authorization(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			Response(w, http.StatusNotFound, Message{Msg: "Not Authorized"})
			return
		}

		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			Response(w, http.StatusUnauthorized, Message{Msg: "Not Authorized"})
			return
		}

		handler.ServeHTTP(w, r)
	})
}
