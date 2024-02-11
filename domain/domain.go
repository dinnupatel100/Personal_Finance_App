package domain

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

type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"categoryname"`
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
