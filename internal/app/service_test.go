package app

import (
	"errors"
	"fmt"
	"testing"

	"github.com/personal-finance-app/db/mocks"
	"github.com/personal-finance-app/domain"
)

// func TestSignup(t *testing.T) {
// 	type args struct {
// 		u User
// 	}
// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				u: User{
// 					FirstName: "sau",
// 					LastName:  "puri",
// 					Email:     "sau@gmail.com",
// 					Password:  "sau@123",
// 				},
// 			},

// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				u: User{
// 					FirstName: "sau",
// 					LastName:  "puri",
// 					Email:     "sau@gmail.com",
// 					Password:  "sau@123",
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Invalid Credential"),
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				u: User{
// 					FirstName: "sau@",
// 					LastName:  "puri",
// 					Email:     "sau@gmail.com",
// 					Password:  "sau@123",
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Invalid Credential"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("Signup", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface)

// 			if err := srv.Signup(tt.args.u); (err != nil) != tt.wantErr {
// 				t.Errorf("error = %v, wantErr %v", err, tt.wantResp)
// 			}
// 		})
// 	}
// }

// func TestLogin(t *testing.T) {
// 	type args struct {
// 		l Login
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				l: Login{
// 					Email:    "sau@gmail.com",
// 					Password: "sau@123",
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				l: Login{
// 					Email:    "sau@gmail",
// 					Password: "sau@123",
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Invalid Login Crede"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("ValidateCredential", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface)

// 			if err := srv.Login(tt.args.l); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 			}

// 		})
// 	}

// }

// func TestSearch(t *testing.T) {

// }

// func TestAddBudget(t *testing.T) {
// 	type args struct {
// 		b Budget
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				b: Budget{
// 					Category:    "food",
// 					Amount:      2000,
// 					StartPeriod: "2024-01-01",
// 					EndPeriod:   "2024-01-02",
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 			// wantResp: domain.Budget{
// 			// 	Category:    "food",
// 			// 	Amount:      2000,
// 			// 	StartPeriod: "2024-01-01",
// 			// 	EndPeriod:   "2024-01-02",
// 			// },
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				b: Budget{
// 					Category:    "food",
// 					StartPeriod: "2024-01-01",
// 					EndPeriod:   "2024-01-02",
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Not found error"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("AddBudget", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface) // this is the actual method

// 			if err := srv.AddBudget(tt.args.b); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 				return
// 			}
// 		})
// 	}
// }

func TestAddCategory(t *testing.T) {
	type args struct {
		c Category
	}

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantResp error
	}{
		{
			name: "success",
			args: args{
				c: Category{
					ID:           1,
					CategoryName: "food",
				},
			},
			wantErr:  false,
			wantResp: nil,
		},
		{
			name: "fail",
			args: args{
				c: Category{
					ID: 1,
				},
			},
			wantErr:  true,
			wantResp: errors.New("Not Found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbInterface := mocks.NewStorer(t)
			dbInterface.On("AddCategory", domain.Category{ID: tt.args.c.ID, CategoryName: tt.args.c.CategoryName}).Return(tt.wantResp)
			srv := NewService(dbInterface) // this is the actual method

			if err := srv.AddCategory(tt.args.c); (err != nil) != tt.wantErr {
				fmt.Println(err)
				t.Errorf("Error: %v want: %v", err, tt.wantResp)
				return
			}
		})
	}
}

// func TestGetAllBudge(t *testing.T) {

// }

// func TestGetBudgetById(t *testing.T) {

// }

// func TestGetTransactionData(t *testing.T) {

// }

// func TestGetBudgetData(t *testing.T) {
// }

// func TestUpdateBudget(t *testing.T) {
// 	type args struct {
// 		b Budget
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				b: Budget{
// 					ID:          1,
// 					Category:    "grocery",
// 					Amount:      2000,
// 					StartPeriod: "2024-01-09",
// 					EndPeriod:   "2024-01-10",
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				b: Budget{
// 					Category:    "grocery",
// 					Amount:      2000,
// 					StartPeriod: "2024-01-01",
// 					EndPeriod:   "2024-01-02",
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Not Found"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("UpdateBudget", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface) // this is the actual method

// 			if err := srv.UpdateBudget(tt.args.b); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 				return
// 			}
// 		})
// 	}

// }

// func TestDeleteBudget(t *testing.T) {
// 	type args struct {
// 		b Budget
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				b: Budget{
// 					ID:          1,
// 					Category:    "grocery",
// 					Amount:      2000,
// 					StartPeriod: "2024-01-09",
// 					EndPeriod:   "2024-01-10",
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				b: Budget{
// 					Category:    "grocery",
// 					Amount:      2000,
// 					StartPeriod: "2024-01-01",
// 					EndPeriod:   "2024-01-02",
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Not Found"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("DeleteBudget", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface) // this is the actual method

// 			if err := srv.DeleteBudget(tt.args.b); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 				return
// 			}
// 		})
// 	}

// }

// func TestAddTransaction(t *testing.T) {
// 	type args struct {
// 		t Transaction
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				t: Transaction{
// 					ID:            1,
// 					Date:          "2020-01-02",
// 					Amount:        200,
// 					Category:      "food",
// 					Tag:           "Vadapav",
// 					Description:   "I ate vadapav",
// 					TransactionID: 12425,
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				t: Transaction{
// 					ID:            1,
// 					Date:          "2020-01-02",
// 					Category:      "food",
// 					Tag:           "Vadapav",
// 					Description:   "I ate vadapav",
// 					TransactionID: 12425,
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Not Found "),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("AddTransaction", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface) // this is the actual method

// 			if err := srv.AddTransaction(tt.args.t); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 				return
// 			}
// 		})
// 	}
// }

// func TestUpdateTransaction(t *testing.T) {
// 	type args struct {
// 		t Transaction
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				t: Transaction{
// 					ID:            1,
// 					Date:          "2020-01-02",
// 					Amount:        200,
// 					Category:      "food",
// 					Tag:           "Vadapav",
// 					Description:   "I ate vadapav",
// 					TransactionID: 12425,
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				t: Transaction{
// 					ID:            1,
// 					Date:          "2020-01-02",
// 					Category:      "food",
// 					Tag:           "Vadapav",
// 					Description:   "I ate vadapav",
// 					TransactionID: 12425,
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Not Found "),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("UpdateTransaction", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface) // this is the actual method

// 			if err := srv.UpdateTransaction(tt.args.t); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 				return
// 			}
// 		})
// 	}

// }

// func TestDeleteTransaction(t *testing.T) {
// 	type args struct {
// 		t Transaction
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantResp error
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				t: Transaction{
// 					ID:            1,
// 					Date:          "2020-01-02",
// 					Amount:        200,
// 					Category:      "food",
// 					Tag:           "Vadapav",
// 					Description:   "I ate vadapav",
// 					TransactionID: 12425,
// 				},
// 			},
// 			wantErr:  false,
// 			wantResp: nil,
// 		},
// 		{
// 			name: "fail",
// 			args: args{
// 				t: Transaction{
// 					ID:            1,
// 					Date:          "2020-01-02",
// 					Category:      "food",
// 					Tag:           "Vadapav",
// 					Description:   "I ate vadapav",
// 					TransactionID: 12425,
// 				},
// 			},
// 			wantErr:  true,
// 			wantResp: errors.New("Not Found "),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbInterface := mocks.NewStorer(t)
// 			dbInterface.On("DeleteTransaction", mock.Anything).Return(tt.wantResp)
// 			srv := NewService(dbInterface) // this is the actual method

// 			if err := srv.DeleteTransaction(tt.args.t); (err != nil) != tt.wantErr {
// 				t.Errorf("Error: %v want: %v", err, tt.wantResp)
// 				return
// 			}
// 		})
// 	}

// }
