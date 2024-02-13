package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/personal-finance-app/internal/app"
	utils "github.com/personal-finance-app/utils/jwt"
	validate "github.com/personal-finance-app/utils/validation"
)

func signup(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user app.User
		err := json.NewDecoder(r.Body).Decode(&user)

		err = validate.ValidateUser(user)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		if err = validate.ValidateName(user.FirstName); err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: NameError})
			return
		}
		if err = validate.ValidateName(user.LastName); err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: NameError})
			return
		}

		if err = validate.ValidateEmail(user.Email); err != nil {
			fmt.Println(err)
			Response(w, http.StatusBadRequest, Message{Msg: EmailError})
			return
		}

		if err = validate.ValidatePassword(user.Password); err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: PasswordError})
			return
		}

		err = service.Signup(user)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: InternalServerError})
			return
		}

		Response(w, http.StatusOK, Message{Msg: Signup})
	}

}

func login(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user app.Login
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		err = validate.ValidateLogin(user)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		if err = validate.ValidateEmail(user.Email); err != nil {

			Response(w, http.StatusBadRequest, Message{Msg: EmailError})
			return
		}

		err = service.Login(user)
		if err != nil {
			Response(w, http.StatusUnauthorized, Message{Msg: CredentialsError})
			return
		}

		token, err := utils.CreateToken(user.Email)
		if err != nil {
			Response(w, http.StatusInternalServerError, InternalServerError)
			return
		}

		Response(w, http.StatusOK, app.Token{TokenString: token, Message: Login})
	}

}
