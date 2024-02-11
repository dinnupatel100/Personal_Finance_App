package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() Storer {
	var err error
	DB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
	return &store{
		db: DB,
	}

}

func createTables() {

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createCategoryTable := `
		CREATE TABLE IF NOT EXISTS category(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			category_name TEXT NOT NULL UNIQUE,
			FOREIGN KEY(id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createCategoryTable)
	if err != nil {
		panic("Could not create users table")
	}

	createTransactionTable := `
		CREATE TABLE IF NOT EXISTS transactions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date DATE NOT NULL,
			amount BIGINT NOT NULL,
			category TEXT NOT NULL,
			tag TEXT NOT NULL,
			description TEXT NOT NULL,
			transaction_id INTEGER UNIQUE,
			FOREIGN KEY(id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createTransactionTable)

	if err != nil {
		panic("Could not create transaction table")
	}

	createBudgetTable := `
		CREATE TABLE IF NOT EXISTS budgets(
			id 	INTEGER PRIMARY KEY AUTOINCREMENT,
			category TEXT NOT NULL,
			amount INTEGER NOT NULL,
			startperiod DATE NOT NULL,
			endperiod DATE NOT NULL,
			FOREIGN KEY(id) REFERENCES transactions(id),
			FOREIGN KEY(id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createBudgetTable)

	if err != nil {
		panic("Could not create budget table")
	}
}
