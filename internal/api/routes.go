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

	// search
	authRoute.HandleFunc("/search", searchTransaction(service)).Methods(http.MethodGet)

	// Transaction
	authRoute.HandleFunc("/addtransaction", addTransaction(service)).Methods("POST")
	authRoute.HandleFunc("/updatetransaction", updateTransaction(service)).Methods(http.MethodPut)
	authRoute.HandleFunc("/deletetransaction", deleteTransaction(service)).Methods("DELETE")
	authRoute.HandleFunc("/getalltransaction", getAllTransactions(service)).Methods("GET")
	authRoute.HandleFunc("/getonetransaction", getTransaction(service)).Methods("GET")

	//Budget
	authRoute.HandleFunc("/addbudget", addBudget(service)).Methods("POST")
	authRoute.HandleFunc("/getallbudget", getAllBudget(service)).Methods("GET")
	authRoute.HandleFunc("/pendingbudget", pendingBudget(service)).Methods("GET")
	authRoute.HandleFunc("/deletebudget", deleteBudget(service)).Methods("DELETE")
	authRoute.HandleFunc("/updatebudget", updateBudget(service)).Methods("PUT")

}
