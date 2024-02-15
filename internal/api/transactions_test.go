package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/personal-finance-app/domain"

	"github.com/personal-finance-app/internal/app"
	"github.com/personal-finance-app/internal/app/mocks"
	"github.com/stretchr/testify/mock"
)

func TestAddTransaction(t *testing.T) {

	// tests := []struct {
	// 	name           string
	// 	args           app.Transaction // input
	// 	expectedOutput error
	// 	wantErr        bool
	// 	setup          func(srvInterface *mocks.Service)
	// }{
	// 	{
	// 		name: "success",
	// 		args: app.Transaction{
	// 			ID:            123,
	// 			Date:          "2024-01-02",
	// 			Amount:        2000,
	// 			Category:      "food",
	// 			Tag:           "restaurant",
	// 			Description:   "Some Description",
	// 			TransactionID: 12344,
	// 		},
	// 		expectedOutput: nil,
	// 		wantErr:        false,
	// 		setup: func(srvInterface *mock.Serv	ice) {

	// 		},
	// 	},
	// }

	transaction := app.Transaction{
		ID:            123,
		Date:          "2024-01-02",
		Amount:        2000,
		Category:      "food",
		Tag:           "restaurant",
		Description:   "Some Description",
		TransactionID: 12344,
	}

	jsonData, err := json.Marshal(transaction)
	if err != nil {
		t.Errorf(err.Error())
	}

	req, err := http.NewRequest("POST", "/api/addtransaction", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()

	mockService := mocks.NewService(t)
	mockService.On("AddTransaction", mock.Anything).Return(nil)
	handler := http.HandlerFunc(addTransaction(mockService))
	handler.ServeHTTP(resp, req)

	statusCode := resp.Code
	if statusCode != http.StatusOK {
		t.Errorf("Got : %v , Want %v ", statusCode, http.StatusOK)
	}

}

func TestUpdateTransaction(t *testing.T) {
	transaction := app.Transaction{
		ID:            12,
		Date:          "2024-01-05",
		Amount:        3000,
		Category:      "shopping",
		Tag:           "cloths",
		Description:   "buy cloths",
		TransactionID: 12345,
	}

	jsonData, err := json.Marshal(transaction)
	if err != nil {
		t.Errorf(err.Error())
	}

	req, err := http.NewRequest("GET", "/api/updatetransaction?id=1", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error())
	}
	resp := httptest.NewRecorder()

	mockService := mocks.NewService(t)
	mockService.On("UpdateTransaction", mock.Anything).Return(nil)
	handler := http.HandlerFunc(updateTransaction(mockService))
	handler.ServeHTTP(resp, req)

	statusCode := resp.Code
	if statusCode != http.StatusOK {
		t.Errorf("Got : %v , Want %v ", statusCode, http.StatusOK)
	}
}

func TestDeleteTransaction(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/deletetransaction?id=1", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()

	mockService := mocks.NewService(t)
	mockService.On("GetTransactionById", int64(1)).Return(&domain.Transaction{}, nil)
	mockService.On("DeleteTransaction", mock.Anything).Return(nil)
	handler := http.HandlerFunc(deleteTransaction(mockService))

	handler.ServeHTTP(resp, req)
	statusCode := resp.Code
	if statusCode != http.StatusOK {
		t.Errorf("Got : %v , Want %v ", statusCode, http.StatusOK)
	}

}

func TestGetAllTransaction(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/getalltransaction", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	mockService := mocks.NewService(t)
	mockService.On("GetAllTransactions").Return([]domain.Transaction{}, nil)
	handler := http.HandlerFunc(getAllTransactions(mockService))

	handler.ServeHTTP(resp, req)
	statusCode := resp.Code
	if statusCode != http.StatusOK {
		t.Errorf("Got : %v , Want %v ", statusCode, http.StatusOK)
	}

}

func TestGetTransaction(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/getonetransaction?category=food", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	mockService := mocks.NewService(t)

	mockService.On("GetTransactionByCategory", mock.Anything).Return([]domain.Transaction{}, nil)
	handler := http.HandlerFunc(getTransaction(mockService))

	handler.ServeHTTP(resp, req)
	statusCode := resp.Code
	if statusCode != http.StatusOK {
		t.Errorf("Got : %v , Want %v ", statusCode, http.StatusOK)
	}

}
