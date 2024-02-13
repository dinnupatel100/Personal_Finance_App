package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/personal-finance-app/internal/app"
	utils "github.com/personal-finance-app/utils/validation"
)

// POST Reuqest
func addTransaction(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var transaction app.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		err = utils.ValidateTransaction(transaction)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		err = service.AddTransaction(transaction)

		if err != nil {
			fmt.Println(err)
			Response(w, http.StatusBadRequest, Message{Msg: UniqueTransaction})
			return
		}

		Response(w, http.StatusBadRequest, Message{Msg: Create})
	}

}

func updateTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}

		var updatedTransaction app.Transaction

		err := json.NewDecoder(r.Body).Decode(&updatedTransaction)
		if err != nil {
			fmt.Println(err)
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		err = utils.ValidateTransaction(updatedTransaction)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		updatedTransaction.ID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err)
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		err = service.UpdateTransaction(updatedTransaction)
		if err != nil {
			fmt.Println(err)
			if err.Error() == NoResourseFound {
				Response(w, http.StatusBadRequest, Message{Msg: NoResourseFound})
				return
			}
			return
		}

		Response(w, http.StatusOK, Message{Msg: Update})
	}

}

func deleteTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")
		if id == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}

		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
		}

		transaction, err := service.GetTransactionById(i)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: NoResourseFound})
			return
		}

		err = service.DeleteTransaction(app.Transaction(*transaction))

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: NoResourseFound})
			return
		}

		Response(w, http.StatusOK, Message{Msg: Delete})
	}

}

func getAllTransactions(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := service.GetAllTransactions()
		if err != nil {
			Response(w, http.StatusBadRequest, FetchingError)
			return
		}
		jsonData, err := json.MarshalIndent(transactions, " ", "\t")
		if err != nil {
			Response(w, http.StatusBadRequest, RequestError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}

}

func getTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")
		if category == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}

		transction, err := service.GetTransactionByCategory(category)
		if err != nil {
			if err.Error() == NoResourseFound {
				Response(w, http.StatusBadRequest, Message{Msg: NoResourseFound})
				return
			}

			Response(w, http.StatusBadRequest, FetchingError)
			return
		}

		jsonData, err := json.MarshalIndent(transction, " ", "\t")
		if err != nil {
			Response(w, http.StatusBadRequest, RequestError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}
