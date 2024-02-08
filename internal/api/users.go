package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/personal-finance-app/internal/app"
	utils "github.com/personal-finance-app/utils/jwt"
)

func signup(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user app.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not pass the request"})
			return
		}

		err = service.Signup(user)

		if err != nil {
			fmt.Println("Err : ", err)
			Response(w, http.StatusBadRequest, Message{Msg: "Could not Signup"})
			return
		}

		Response(w, http.StatusOK, Message{Msg: "Signup successsfully"})
	}

}

func login(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user app.Login
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			fmt.Println("Error :", err)
			Response(w, http.StatusBadRequest, Message{Msg: "Could not parse the request"})
			return
		}

		err = service.Login(user) // check the user credentials
		if err != nil {
			Response(w, http.StatusUnauthorized, Message{Msg: "Invalid Credentials"})
			return
		}

		token, err := utils.CreateToken(user.Email)
		if err != nil {
			Response(w, http.StatusInternalServerError, err.Error())
		}

		Response(w, http.StatusOK, app.Token{TokenString: token, Message: "Login Successfully"})
	}

}
