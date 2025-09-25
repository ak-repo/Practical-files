package jwtpkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(secretKey string, userID uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Minute * 10).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}

func ValidateToken(secretKey, tokenSTR string) error {

	token, err := jwt.Parse(tokenSTR, func(token *jwt.Token) (interface{}, error) {

		// method verification
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("method not ")

		}

		return []byte(secretKey), nil
	})

	if !token.Valid || err != nil {
		return err
	}

	return nil

}
