package utils

import (
	"errors"
	"strings"
	"time"
	"unicode"

	"github.com/personal-finance-app/internal/app"
)

func ValidateName(name string) error {
	if len(name) == 0 {
		return errors.New("Name can't be empty")
	}

	if len(name) < 1 && len(name) > 10 {
		return errors.New("Check the length of name")
	}

	for _, value := range name {
		if value == ' ' || unicode.IsPunct(value) {
			return errors.New("Name contains the special characters")
		}
	}

	return nil
}

func checkValidCharecter(firstPart string) bool {
	if len(firstPart) == 0 {
		return false
	}
	for _, char := range firstPart {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || strings.ContainsAny(string(char), "!#$%&'*+-/=?^_`{|}~.")) {
			return false
		}
	}
	return true
}

func checkSecondPart(secondPart string) bool {
	if len(secondPart) == 0 {
		return false
	}

	subdomains := strings.Split(secondPart, ".")
	if len(subdomains) < 2 {
		return false
	}

	for _, subDomain := range subdomains {
		if !checkSubDomain(subDomain) {
			return false
		}
	}

	return true
}

func checkSubDomain(subdomain string) bool {
	if len(subdomain) == 0 {
		return false
	}

	for _, char := range subdomain {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '-' {
			return false
		}
	}

	return true
}

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("Email can not be empty")
	}

	index := strings.Index(email, "@")
	if index <= 0 || index == len(email)-1 {
		return errors.New("Invalid email")
	}

	firstPart := email[:index]
	secondPart := email[index+1:]

	if !checkValidCharecter(firstPart) {
		return errors.New("Contains invalid character")
	}

	if !checkSecondPart(secondPart) {
		return errors.New("Invalid domain part")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 0 {
		return errors.New("Password can not be weak")
	}

	var (
		checkLowerCase   bool
		checkDigit       bool
		checkSpecialChar bool
	)

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			checkLowerCase = true
		case unicode.IsDigit(char):
			checkDigit = true
		case unicode.IsPunct(char):
			checkSpecialChar = true
		}
	}

	if !checkLowerCase || !checkDigit || !checkSpecialChar {
		return errors.New("Password seems to be week")
	}

	return nil
}

func ValidateDate(date string) error {
	dateFormat := "2006-01-02"
	_, err := time.Parse(dateFormat, date)
	if err != nil {
		return errors.New("Invalid Date")
	}
	return nil
}

func ValidateUser(u app.User) error {
	if len(u.FirstName) <= 0 {
		return errors.New("First Name can not be empty")
	}

	if len(u.LastName) <= 0 {
		return errors.New("Last name can not be empty")
	}

	if len(u.Email) <= 0 {
		return errors.New("Email can not be empty")
	}

	if len(u.Password) <= 0 {
		return errors.New("Password can not be empty")
	}

	if len(u.Password) < 5 {
		return errors.New("Password is too short")
	}
	return nil
}

func ValidateLogin(u app.Login) error {
	if len(u.Email) <= 0 {
		return errors.New("Email can not be empty")
	}

	if len(u.Password) <= 0 {
		return errors.New("Password can not be empty")
	}

	return nil
}

func ValidateBudget(b app.Budget) error {
	if b.Amount <= 0 {
		return errors.New("Amount can not be empty")
	}

	if len(b.Category) <= 0 {
		return errors.New("Category cannot be empty")
	}

	if ValidateName(b.Category) != nil {
		return errors.New("Please provide valid category name")
	}

	if len(b.StartPeriod) <= 0 {
		return errors.New("Start Period cannot be empty")
	}
	if ValidateDate(b.StartPeriod) != nil {
		return errors.New("Invalid date")
	}
	if len(b.EndPeriod) <= 0 {
		return errors.New("End Period cannot be empty")
	}
	if ValidateDate(b.EndPeriod) != nil {
		return errors.New("Invalid date")
	}

	return nil
}

func ValidateTransaction(t app.Transaction) error {
	if len(t.Date) <= 0 {
		return errors.New("Date can not be empty")
	}

	if ValidateDate(t.Date) != nil {
		return errors.New("Invalid Date")
	}

	if t.Amount <= 0 {
		return errors.New("Amount can not be empty")
	}

	if len(t.Category) <= 0 {
		return errors.New("Category can not be empty")
	}

	if ValidateName(t.Category) != nil {
		return errors.New("Invalid category name")
	}

	if ValidateName(t.Tag) != nil {
		return errors.New("Invalid tag name")
	}

	if len(t.Tag) <= 0 {
		return errors.New("Tag can not be empty")
	}

	if len(t.Description) <= 0 {
		return errors.New("Description can not be empty")
	}

	if t.TransactionID <= 0 {
		return errors.New("Transaction ID can not be empty")
	}

	return nil
}

func ValidateCatgory(c app.Category) error {
	if len(c.CategoryName) <= 0 {
		return errors.New("Category can not be empty")
	}

	if ValidateName(c.CategoryName) != nil {
		return errors.New("Category must be valid")
	}
	return nil
}
