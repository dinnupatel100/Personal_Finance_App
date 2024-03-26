package app

import (
	"errors"
	"time"

	"github.com/personal-finance-app/domain"
	"github.com/personal-finance-app/internal/db"
)

func NewService(str db.Storer) Service {
	return &service{
		store: str,
	}
}

func (s *service) Signup(u User) error {
	user := domain.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}

	return s.store.Signup(user)
}

func (s *service) Login(l Login) error {
	login := Login{
		ID:       l.ID,
		Email:    l.Email,
		Password: l.Password,
	}
	return s.store.ValidateCredential(domain.Login(login))
}

func (s *service) Search(tag string) ([]domain.Transaction, error) {
	return s.store.Search(tag)
}

func (s *service) AddBudget(b Budget) error {

	addBudget := domain.Budget{
		ID:          b.ID,
		Category:    b.Category,
		Amount:      b.Amount,
		StartPeriod: b.StartPeriod,
		EndPeriod:   b.EndPeriod,
	}

	return s.store.AddBudget(addBudget)
}

func (s *service) AddCategory(c Category) error {
	addCategory := domain.Category{
		ID:           c.ID,
		CategoryName: c.CategoryName,
	}

	return s.store.AddCategory(addCategory)
}

func (s *service) GetAllBudgets() ([]domain.Budget, error) {
	return s.store.GetAllBudgets()
}

func (s *service) DeleteBudget(b Budget) error {
	deleteBudget := domain.Budget{
		ID:          b.ID,
		Category:    b.Category,
		Amount:      b.Amount,
		StartPeriod: b.StartPeriod,
		EndPeriod:   b.EndPeriod,
	}

	return s.store.DeleteBudget(deleteBudget)
}

func (s *service) GetBudgetById(id int64) (*domain.Budget, error) {
	return s.store.GetBudgetById(id)
}

func (s *service) UpdateBudget(b Budget) error {
	updateBudget := domain.Budget{
		ID:          b.ID,
		Category:    b.Category,
		Amount:      b.Amount,
		StartPeriod: b.StartPeriod,
		EndPeriod:   b.EndPeriod,
	}

	return s.store.UpdateBudget(updateBudget)
}

func (s *service) AddTransaction(t Transaction) error {
	addTransaction := domain.Transaction{
		ID:            t.ID,
		Date:          t.Date,
		Amount:        t.Amount,
		Category:      t.Category,
		Tag:           t.Tag,
		Description:   t.Description,
		TransactionID: t.TransactionID,
	}

	return s.store.AddTransaction(addTransaction)

}

func (s *service) UpdateTransaction(t Transaction) error {
	updateTransaction := domain.Transaction{
		ID:            t.ID,
		Date:          t.Date,
		Amount:        t.Amount,
		Category:      t.Category,
		Tag:           t.Tag,
		Description:   t.Description,
		TransactionID: t.TransactionID,
	}

	return s.store.UpdateTransaction(updateTransaction)
}

func (s *service) DeleteTransaction(t Transaction) error {
	deleteTransaction := domain.Transaction{
		ID:            t.ID,
		Date:          t.Date,
		Amount:        t.Amount,
		Category:      t.Category,
		Description:   t.Description,
		TransactionID: t.TransactionID,
	}

	return s.store.DeleteTransaction(deleteTransaction)
}

func (s *service) GetTransactionById(id int64) (*domain.Transaction, error) {
	return s.store.GetTransactionById(id)
}

func (s *service) GetAllTransactions() ([]domain.Transaction, error) {
	return s.store.GetAllTransactions()
}

func (s *service) GetTransactionByCategory(category string) ([]domain.Transaction, error) {
	return s.store.GetTransactionByCategory(category)
}

func (s *service) GetPendingAmount(category string) (int64, error) {

	transactionData, err := s.store.GetTotalTransactionBYCategory(category)
	if err != nil {
		return 0, errors.New("Could not find the transaction data : ")
	}

	budgetData, err := s.store.GetTotalBudgetByCategory(category)
	if err != nil {
		return 0, errors.New("Could not get the budget data")
	}

	pendingAmount := budgetData[category] - transactionData[category]

	return pendingAmount, nil
}

func (s *service) GetTransactionByDate(startDate, endDate time.Time) ([]domain.Transaction, error) {
	return s.store.GetTransactionFromTo(startDate, endDate)
}


func (s *service) GetAllCategory() ([]domain.Category, error) {
	return s.store.GetAllCategory()
}


func (s *service) GetCategoryById(id int64) (*domain.Category, error) {
	return s.store.GetCategoryById(id)
}


func (s *service) DeleteCategory(c Category) error {
	deleteCategory := domain.Category{
		ID:            c.ID,
		CategoryName:   c.CategoryName,
	}

	return s.store.DeleteCategory(deleteCategory)
}