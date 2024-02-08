package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/personal-finance-app/internal/app"
)

func addBudget(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var budget app.Budget

		err := json.NewDecoder(r.Body).Decode(&budget)

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		err = service.AddBudget(budget)

		if err != nil {
			fmt.Println("Error :", err)
			Response(w, http.StatusInternalServerError, Message{Msg: "Could not create the budget some erorr occured"})
			return
		}

		Response(w, http.StatusOK, Message{Msg: "Budget Add Successfully"})
	}

}

func getAllBudget(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		budgets, err := service.GetAllBudgets()
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusOK, budgets)
	}

}

func pendingBudget(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category") // query parameter

		var pendingFoodAmount, pendingGroceryAmount, pendingShoppingAmount float64
		transactionData, err := service.GetTransactionData()

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		budgetData, err := service.GetBudgetData()

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		for index, value := range budgetData {

			if val, ok := transactionData[index]; ok {
				if index == "food" {
					pendingFoodAmount = float64(value - val)

				} else if index == "shopping" {
					pendingShoppingAmount = float64(value - val)

				} else if index == "grocery" {
					pendingGroceryAmount = float64(value - val)
				}
			}

		}

		// for the food
		if pendingFoodAmount < 0 {

			if category == "food" {
				if pendingFoodAmount < 0 {
					value := strconv.FormatFloat(math.Abs(pendingFoodAmount), 'f', -1, 64)
					w.Write([]byte("Your payment is exceed for food by : "))
					w.Write([]byte(value))
				} else {
					value := strconv.FormatFloat(math.Abs(pendingFoodAmount), 'f', -1, 64)
					w.Write([]byte("Your pending amount for food : "))
					w.Write([]byte(value))

				}
			}

			// for the grocery
			if category == "grocery" {
				if pendingGroceryAmount < 0 {
					value := strconv.FormatFloat(math.Abs(pendingGroceryAmount), 'f', -1, 64)
					w.Write([]byte("Your payment is exceed for grocery by : "))
					w.Write([]byte(value))
				} else {
					value := strconv.FormatFloat(math.Abs(pendingGroceryAmount), 'f', -1, 64)
					w.Write([]byte("Your pending amount for grocery : "))
					w.Write([]byte(value))
				}
			}

			// for the shopping
			if category == "shopping" {
				if pendingShoppingAmount < 0 {
					value := strconv.FormatFloat(math.Abs(pendingShoppingAmount), 'f', -1, 64)
					w.Write([]byte("Your payment is exceed for shopping by  : "))
					w.Write([]byte(value))
				} else {
					value := strconv.FormatFloat(math.Abs(pendingShoppingAmount), 'f', -1, 64)
					w.Write([]byte("Your pending amount for shopping : "))
					w.Write([]byte(value))
				}
			}
		}
	}
	// three categories : Food , Grocery ,  Shopping

}

func deleteBudget(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := r.URL.Query().Get("id") // query parameter

		// UserID := context.GetInt64("id") // it gives the different id

		i, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Cannot convert into int"})
		}

		budget, err := service.GetBudgetById(i)

		// if err != nil {
		// 	Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
		// 	return
		// }

		err = service.DeleteBudget(app.Budget(*budget))

		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusBadRequest, Message{Msg: "Budget delete successful"})
	}

}

func updateBudget(service app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := r.URL.Query().Get("id") // query parameter

		// userId := context.GetInt64("budgetId")
		i, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Cannot convert into int"})
		}

		// budget, err := service.GetBudgetById(i)
		// if err != nil {
		// 	Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
		// 	return
		// }

		// if budget.ID != userId {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to update budget you tried with different budget"})
		// 	return
		// }

		var updateBudget app.Budget

		err = json.NewDecoder(r.Body).Decode(&updateBudget)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		updateBudget.ID = i
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: "Cannot convert the string"})
		}
		err = service.UpdateBudget(updateBudget)
		if err != nil {
			fmt.Println("Error is :", err)
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusBadRequest, Message{Msg: "Update budget successfully"})

	}

}
