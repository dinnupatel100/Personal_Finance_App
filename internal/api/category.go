package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/personal-finance-app/internal/app"
)

func addCategory(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var category app.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			fmt.Println("Err", err)
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		err = service.AddCategory(category)
		if err != nil {
			fmt.Println("Err", err)
			Response(w, http.StatusBadRequest, RequestError)
			return
		}

		Response(w, http.StatusOK, Message{Msg: Create})
	}
}
