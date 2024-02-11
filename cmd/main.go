package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/personal-finance-app/db"
	"github.com/personal-finance-app/internal/api"
	"github.com/personal-finance-app/internal/app"
)

func main() {
	mux := mux.NewRouter()
	interfaceGotFromDatabase := db.InitDB()
	interfaceGotFromServiceLayer := app.NewService(interfaceGotFromDatabase)

	api.RegisteredRoutes(mux, interfaceGotFromServiceLayer)

	fmt.Println("Server Started...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
