package api

import (
	"fmt"
	"net/http"

	"github.com/personal-finance-app/internal/app"
)

func searchTransaction(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := r.URL.Query().Get("tag")
		if tag == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}
		searchTransaction, err := service.Search(tag)
		if err != nil {
			fmt.Println(err)
			if err.Error() == NoResourseFound {
				Response(w, http.StatusNotFound, Message{Msg: NoResourseFound})
				return
			}
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		Response(w, http.StatusOK, searchTransaction)
	}
}
