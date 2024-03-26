package app

import (
	"time"

	"github.com/personal-finance-app/domain"
	"github.com/personal-finance-app/internal/db"
)

type Service interface {
	Signup(User) error
	Login(Login) error

	Search(string) ([]domain.Transaction, error)
	AddCategory(Category) error
	GetCategoryById(int64) (*domain.Category, error)
	DeleteCategory(Category) error

	AddBudget(Budget) error
	GetAllBudgets() ([]domain.Budget, error)
	DeleteBudget(Budget) error
	GetBudgetById(int64) (*domain.Budget, error)
	UpdateBudget(Budget) error
	AddTransaction(Transaction) error
	UpdateTransaction(Transaction) error
	DeleteTransaction(Transaction) error
	GetTransactionById(id int64) (*domain.Transaction, error)
	GetAllTransactions() ([]domain.Transaction, error)
	GetTransactionByCategory(string) ([]domain.Transaction, error)

	GetPendingAmount(string) (int64, error)
	GetTransactionByDate(time.Time, time.Time) ([]domain.Transaction, error)
	GetAllCategory() ([]domain.Category, error)
}

type service struct {
	store db.Storer
}

type User struct {
	ID        int64
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required"`
}

type Login struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" json:",omitempty" binding:"required"`
}

type Budget struct {
	ID          int64
	Category    string `json:"category" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	StartPeriod string `json:"startperiod" binding:"required"`
	EndPeriod   string `json:"endperiod" binding:"required"`
}

type Transaction struct {
	ID            int64  `json:"id"`
	Date          string `json:"date" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Category      string `json:"category" binding:"required"`
	Tag           string `json:"tag" binding:"required"`
	Description   string `json:"description" binding:"required"`
	TransactionID int64  `json:"transaction_id"`
}

type Token struct {
	TokenString string `json:"token"`
	Message     string `json:"message"`
}

type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"categoryname"`
}
