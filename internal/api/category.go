package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/personal-finance-app/internal/app"
	utils "github.com/personal-finance-app/utils/validation"
)

func addCategory(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var category app.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			fmt.Println("Err", err)
			Response(w, http.StatusInternalServerError, Message{Msg: InternalServerError})
			return
		}

		err = utils.ValidateCatgory(category)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		err = service.AddCategory(category)
		if err != nil {
			fmt.Println("Err", err)
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		Response(w, http.StatusOK, Message{Msg: Create})
	}
}


func getAllCategory(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := service.GetAllCategory()
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: FetchingError})
			return
		}
		jsonData, err := json.MarshalIndent(categories, " ", "\t")
		if err != nil {
			Response(w, http.StatusInternalServerError, Message{Msg: InternalServerError})
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}


func deleteCategory(service app.Service) func(w http.ResponseWriter, h *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := r.URL.Query().Get("id")
		if paramId == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}

		i, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		category, err := service.GetCategoryById(i)
		if err != nil {
			if err.Error() == NoResourseFound {
				Response(w, http.StatusNotFound, Message{Msg: NoResourseFound})
				return
			}
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		err = service.DeleteCategory(app.Category(*category))
		if err != nil {
			fmt.Println(err)
			Response(w, http.StatusInternalServerError, Message{Msg: InternalServerError})
			return
		}

		Response(w, http.StatusOK, Message{Msg: Delete})
	}

}

