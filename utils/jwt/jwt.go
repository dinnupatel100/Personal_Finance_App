package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email string) (string, error) {
	var err error
	getClaims := jwt.MapClaims{} // this is used to store the payloads
	getClaims["authorised"] = true
	getClaims["email"] = email
	getClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, getClaims)
	tokenString, err := token.SignedString([]byte("secrete-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// get the signing method
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid signning method")
		}

		return []byte("secrete-key"), nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
