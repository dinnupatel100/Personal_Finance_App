package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/personal-finance-app/internal/app"
	"github.com/personal-finance-app/middleware"
)

func RegisteredRoutes(router *mux.Router, service app.Service) {

	// User
	router.HandleFunc("/signup", signup(service)).Methods(http.MethodPost)
	router.HandleFunc("/login", login(service)).Methods(http.MethodPost)

	authRoute := router.PathPrefix("/api").Subrouter()
	authRoute.Use(middleware.Authorization)

	// category
	authRoute.HandleFunc("/addcategory", addCategory(service)).Methods(http.MethodPost)
	authRoute.HandleFunc("/getallcategory", getAllCategory(service)).Methods(http.MethodGet)

	// search
	authRoute.HandleFunc("/search", searchTransaction(service)).Methods(http.MethodGet)

	// Transaction
	authRoute.HandleFunc("/addtransaction", addTransaction(service)).Methods(http.MethodPost)
	authRoute.HandleFunc("/updatetransaction", updateTransaction(service)).Methods(http.MethodPut)
	authRoute.HandleFunc("/deletetransaction", deleteTransaction(service)).Methods(http.MethodDelete)
	authRoute.HandleFunc("/getalltransaction", getAllTransactions(service)).Methods(http.MethodGet)
	authRoute.HandleFunc("/getonetransaction", getTransaction(service)).Methods(http.MethodGet)

	//Budget
	authRoute.HandleFunc("/addbudget", addBudget(service)).Methods(http.MethodPost)
	authRoute.HandleFunc("/getallbudget", getAllBudget(service)).Methods(http.MethodGet)
	authRoute.HandleFunc("/pendingbudget", pendingBudget(service)).Methods(http.MethodGet)
	authRoute.HandleFunc("/deletebudget", deleteBudget(service)).Methods(http.MethodDelete)
	authRoute.HandleFunc("/updatebudget", updateBudget(service)).Methods(http.MethodPut)

	authRoute.HandleFunc("/from-to", getTransactionByDate(service)).Methods(http.MethodGet)
}
