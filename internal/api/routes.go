package api

import (
	"github.com/gorilla/mux"
	"github.com/personal-finance-app/internal/app"
)

func RegisteredRoutes(router *mux.Router, service app.Service) {

	// User
	router.HandleFunc("/signup", signup(service)).Methods("POST")
	router.HandleFunc("/login", login(service)).Methods("POST")

	// Transaction
	router.HandleFunc("/addtransaction", addTransaction(service)).Methods("POST")
	router.HandleFunc("/updatetransaction", updateTransaction(service)).Methods("PUT")
	router.HandleFunc("/deletetransaction", deleteTransaction(service)).Methods("DELETE")
	router.HandleFunc("/getalltransaction", getAllTransactions(service)).Methods("GET")
	router.HandleFunc("/getonetransaction", getTransaction(service)).Methods("GET")

	//Budget
	router.HandleFunc("/addbudget", addBudget(service)).Methods("POST")
	router.HandleFunc("/getallbudget", getAllBudget(service)).Methods("GET")
	router.HandleFunc("/pendingbudget", pendingBudget(service)).Methods("GET")
	router.HandleFunc("/deletebudget", deleteBudget(service)).Methods("DELETE")
	router.HandleFunc("/updatebudget", updateBudget(service)).Methods("PUT")

}
f