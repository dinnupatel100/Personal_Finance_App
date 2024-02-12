package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/personal-finance-app/internal/app"
)

func addBudget(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var budget app.Budget

		err := json.NewDecoder(r.Body).Decode(&budget)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		err = service.AddBudget(budget)

		if err != nil {
			fmt.Println("Error :", err)
			Response(w, http.StatusInternalServerError, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusOK, Message{Msg: Create})
	}

}

func getAllBudget(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		budgets, err := service.GetAllBudgets()
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusOK, budgets)
	}

}

func pendingBudget(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")
		if category == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}

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

				if index == "food" && category == "food" {
					pendingFoodAmount = float64(value - val)
					if pendingFoodAmount < 0 {
						value := strconv.FormatFloat(math.Abs(pendingFoodAmount), 'f', -1, 64)
						w.Write([]byte("Your payment is exceed for food by : "))
						w.Write([]byte(value))
					} else {
						value := strconv.FormatFloat(math.Abs(pendingFoodAmount), 'f', -1, 64)
						w.Write([]byte("Your pending amount for food : "))
						w.Write([]byte(value))

					}

				} else if index == "grocery" && category == "grocery" {
					pendingShoppingAmount = float64(value - val)
					if pendingGroceryAmount < 0 {
						value := strconv.FormatFloat(math.Abs(pendingGroceryAmount), 'f', -1, 64)
						w.Write([]byte("Your payment is exceed for grocery by : "))
						w.Write([]byte(value))
					} else {
						value := strconv.FormatFloat(pendingGroceryAmount, 'f', -1, 64)
						fmt.Println(pendingGroceryAmount)
						w.Write([]byte("Your pending amount for grocery : "))
						w.Write([]byte(value))
					}

				} else if index == "shopping" && category == "shopping" {
					pendingGroceryAmount = float64(value - val)
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
			} else {
				Response(w, http.StatusNotFound, Message{Msg: NoResourseFound})
				return
			}

		}
	}
}

func deleteBudget(service app.Service) func(w http.ResponseWriter, h *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := r.URL.Query().Get("id")

		i, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
		}

		budget, err := service.GetBudgetById(i)

		if err != nil {
			fmt.Println(err)
			Response(w, http.StatusBadRequest, Message{Msg: NoResourseFound})
			return
		}

		err = service.DeleteBudget(app.Budget(*budget))

		if err != nil {
			if err.Error() == NoResourseFound {
				Response(w, http.StatusNotFound, Message{Msg: NoResourseFound})
				return
			}
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusBadRequest, Message{Msg: "Budget delete successful"})
	}

}

func updateBudget(service app.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := r.URL.Query().Get("id")
		if paramId == "" {
			Response(w, http.StatusNotFound, Message{Msg: QueryNotFoundError})
			return
		}

		i, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			fmt.Print(err)
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}

		var updateBudget app.Budget

		err = json.NewDecoder(r.Body).Decode(&updateBudget)
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		updateBudget.ID = i
		if err != nil {
			Response(w, http.StatusBadRequest, Message{Msg: RequestError})
			return
		}
		err = service.UpdateBudget(updateBudget)
		if err != nil {
			fmt.Println("Error is :", err)
			Response(w, http.StatusBadRequest, Message{Msg: err.Error()})
			return
		}

		Response(w, http.StatusBadRequest, Message{Msg: Update})

	}

}
