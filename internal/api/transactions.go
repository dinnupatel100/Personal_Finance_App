package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/personal-finance-app/internal/app"
)

// POST Reuqest
func addTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction app.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)

		if err != nil {
			fmt.Println("err", err)
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		err = service.AddTransaction(transaction)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusBadRequest, Message{Msg: "Transaction created successfull"})
	}

	// w.Write([]byte(transaction))

}

func updateTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id") // query parameter

		// transactionId := context.GetInt64("transactionId")
		// transaction, err := app.GetTransactionById(value)

		// if err != nil {
		// 	Response(w , http.StatusBadRequest , Message{Msg: err.Error()})
		// 	return
		// }

		// if transaction.TransactionID != transactionId {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to update transction you tried with different event"})
		// 	return
		// }

		var updatedTransaction app.Transaction

		err := json.NewDecoder(r.Body).Decode(&updatedTransaction)
		if err != nil {
			fmt.Println(err)
			Response(w, http.StatusBadRequest, Message{Msg: "Could not parse the request data into json"})
			return
		}

		updatedTransaction.ID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Cannot convert the string"})
		}

		err = service.UpdateTransaction(updatedTransaction)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not update the transaction"})
			return
		}

		Response(w, http.StatusOK, Message{Msg: "Update Transaction successfully"})
	}

}

func deleteTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id") // query parameter

		// transactionId := context.GetInt64("transactionId")
		i, err := strconv.ParseInt(id, 10, 64)
		transaction, err := service.GetTransactionById(i)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not fetch the transaction"})
			return
		}

		// if transaction.TransactionID != transactionId {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to update transction you tried with different event"})
		// 	return
		// }

		err = service.DeleteTransaction(app.Transaction(*transaction))

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not delete transaction"})
			return
		}

		Response(w, http.StatusOK, Message{Msg: "Transaction deleted successfully"})
	}

}

func getAllTransactions(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := service.GetAllTransactions()
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not fetch the transaction"})
			return
		}
		jsonData, err := json.MarshalIndent(transactions, " ", "\t")
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not convert into json"})
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}

}

func getTransaction(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// paramStr, ok := context.Params.Get("category")
		category := r.URL.Query().Get("category") // query parameter

		transction, err := service.GetTransactionByCategory(category)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not fetch the transaction"})
			return
		}

		jsonData, err := json.MarshalIndent(transction, " ", "\t")
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Could not convert into the json"})
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}
