package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/personal-finance-app/domain"
	utils "github.com/personal-finance-app/utils/hash"
)

var budgets = []domain.Budget{}

type Storer interface {
	Signup(domain.User) error
	ValidateCredential(domain.Login) error

	Search(string) ([]domain.Transaction, error)

	AddCategory(domain.Category) error

	AddBudget(budget domain.Budget) error
	GetAllBudgets() ([]domain.Budget, error)
	DeleteBudget(domain.Budget) error
	GetBudgetById(int64) (*domain.Budget, error)
	GetTotalTransactionBYCategory(string) (map[string]int64, error)
	GetTotalBudgetByCategory(string) (map[string]int64, error)
	UpdateBudget(domain.Budget) error

	AddTransaction(domain.Transaction) error
	UpdateTransaction(t domain.Transaction) error
	DeleteTransaction(domain.Transaction) error
	GetTransactionById(id int64) (*domain.Transaction, error)
	GetAllTransactions() ([]domain.Transaction, error)
	GetTransactionByCategory(string) ([]domain.Transaction, error)
}

type store struct {
	db *sql.DB
}

func (s *store) Signup(u domain.User) error {
	query := `INSERT INTO users(first_name , last_name , email , password) VALUES (?,?,?,?)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, hashPassword)
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if num != 1 || err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (s *store) ValidateCredential(u domain.Login) error {
	query := "SELECT password FROM users WHERE email = ?"

	row := s.db.QueryRow(query, u.Email)
	var retreivedPassword string

	err := row.Scan(&retreivedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(NoResourseFound)
		}
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, retreivedPassword)

	if !passwordIsValid {
		return errors.New("credential invalid")
	}
	return nil
}

func (s *store) Search(tag string) ([]domain.Transaction, error) {
	query := "SELECT * FROM transactions WHERE tag = ?"
	rows, err := s.db.Query(query, tag)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []domain.Transaction

	for rows.Next() {
		var transaction domain.Transaction

		err := rows.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.Category, &transaction.Tag, &transaction.Description, &transaction.TransactionID)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return nil, errors.New(NoResourseFound)
	}

	return transactions, nil

}

func (s *store) AddCategory(c domain.Category) error {
	query := `INSERT INTO category(category_name) VALUES(?)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(c.CategoryName)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return err
	}

	_, err = result.LastInsertId()
	//	c.ID = id
	return err

}

func (s *store) AddBudget(b domain.Budget) error {
	query := "SELECT COUNT(*) FROM budgets WHERE category = ?"
	var count int

	row := s.db.QueryRow(query, b.Category)
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(NoResourseFound)
		}
		return err
	}

	if count > 0 {
		return errors.New("category already exists")
	}

	query = "SELECT COUNT(*) FROM category WHERE category_name = ?"

	row = s.db.QueryRow(query, b.Category)
	err = row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(NoResourseFound)
		}
		return err
	}

	if count == 0 {
		return errors.New("category does not exists in the category")
	}

	query = `INSERT INTO budgets(category , amount , startperiod , endperiod)
	VALUES(?,?,?,?)`

	stmt, err := s.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(b.Category, b.Amount, b.StartPeriod, b.EndPeriod)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id
	return err
}

func (s *store) GetAllBudgets() ([]domain.Budget, error) {
	query := "SELECT * FROM budgets"
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var budgets []domain.Budget
	for rows.Next() {
		var budget domain.Budget
		err := rows.Scan(&budget.ID, &budget.Category, &budget.Amount, &budget.StartPeriod, &budget.EndPeriod)
		if err != nil {
			if err == sql.ErrNoRows {
				return []domain.Budget{}, errors.New(NoResourseFound)
			}
			return nil, err
		}
		budgets = append(budgets, budget)
	}
	return budgets, nil
}

func (s *store) DeleteBudget(b domain.Budget) error {
	query := "DELETE FROM budgets WHERE id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(b.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		if err == sql.ErrNoRows {
			return errors.New(NoResourseFound)
		}
		return err
	}
	return nil
}

func (s *store) GetBudgetById(id int64) (*domain.Budget, error) {
	query := "SELECT * FROM budgets WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var budget domain.Budget

	err := row.Scan(&budget.ID, &budget.Category, &budget.Amount, &budget.StartPeriod, &budget.EndPeriod)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(NoResourseFound)
		}
		return nil, err
	}

	return &budget, nil
}

func (s *store) GetTotalTransactionBYCategory(category string) (map[string]int64, error) {
	tMap := make(map[string]int64)
	query := `SELECT amount FROM transactions WHERE category = ?`
	table, err := s.db.Query(query, category)

	if err != nil {
		return nil, errors.New(InternalServerError)
	}

	defer table.Close()

	for table.Next() {
		var amount string
		err := table.Scan(&amount)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New(NoResourseFound)
			}
			return nil, errors.New(InternalServerError)
		}

		amountValue, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			return nil, errors.New(NoResourseFound)
		}
		tMap[category] = tMap[category] + amountValue
	}

	return tMap, nil
}

func (s *store) GetTotalBudgetByCategory(category string) (map[string]int64, error) {
	bMap := make(map[string]int64)
	query := `SELECT amount FROM budgets WHERE category = ? `
	table, err := s.db.Query(query, category)
	if err != nil {
		return nil, errors.New(InternalServerError)
	}

	defer table.Close()

	for table.Next() {
		var amount string
		err := table.Scan(&amount)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New(NoResourseFound)
			}
			return nil, errors.New(InternalServerError)
		}

		amountValue, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			return nil, errors.New(InternalServerError)
		}
		bMap[category] = bMap[category] + amountValue
	}

	return bMap, nil
}

func (s *store) UpdateBudget(b domain.Budget) error {

	checkQuery := "SELECT COUNT(*) FROM budgets WHERE id = ?"
	checkRow := s.db.QueryRow(checkQuery, b.ID)
	var count int
	if err := checkRow.Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return errors.New(NoResourseFound)
	}

	query := `UPDATE budgets 
		SET category = ?,amount = ? , startperiod = ? , endperiod = ?
		WHERE id = ? 
	`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(b.Category, b.Amount, b.StartPeriod, b.EndPeriod, b.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New(NoResourseFound)
		}
		return err
	}
	return nil
}

func (s *store) AddTransaction(t domain.Transaction) error {
	query := `INSERT INTO transactions(date , amount , category ,tag ,  description , transaction_id)
	VALUES(?,?,?,?,?,?)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(t.Date, t.Amount, t.Category, t.Tag, t.Description, t.TransactionID)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}

func (s *store) UpdateTransaction(t domain.Transaction) error {
	query := `UPDATE transactions
		SET date =?,amount=?,category=?,tag = ?,description=?,transaction_id=?
		WHERE id = ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(t.Date, t.Amount, t.Category, t.Tag, t.Description, t.TransactionID, t.ID)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New(NoResourseFound)
	}

	return nil
}

func (s *store) DeleteTransaction(t domain.Transaction) error {
	query := "DELETE FROM transactions WHERE id = ?"
	stmt, err := s.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(t.ID)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New(NoResourseFound)
	}
	return nil
}

func (s *store) GetTransactionById(id int64) (*domain.Transaction, error) {
	query := "SELECT * FROM transactions WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var transaction domain.Transaction

	err := row.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.Category, &transaction.Tag, &transaction.Description, &transaction.TransactionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(NoResourseFound)
		}
		return nil, err
	}

	return &transaction, nil
}

func (s *store) GetAllTransactions() ([]domain.Transaction, error) {
	query := "SELECT * FROM transactions"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []domain.Transaction

	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.Category, &transaction.Tag, &transaction.Description, &transaction.TransactionID)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New(NoResourseFound)
			}
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil

}

func (s *store) GetTransactionByCategory(category string) ([]domain.Transaction, error) {
	query := `SELECT * FROM transactions WHERE category = ?`
	row, err := s.db.Query(query, category)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer row.Close()

	var transactions []domain.Transaction

	for row.Next() {
		var transaction domain.Transaction
		err := row.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.Category, &transaction.Tag, &transaction.Description, &transaction.TransactionID)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return nil, errors.New(NoResourseFound)
	}
	return transactions, nil
}
