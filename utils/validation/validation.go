package utils

import (
	"errors"
	"strings"
	"unicode"
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
