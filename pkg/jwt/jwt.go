package jwt

import (
	"errors"

	"api.default.indicoinnovation.pt/pkg/crypt"
	"github.com/golang-jwt/jwt"
)

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

func Verify(tokenString string) (bool, error) {
	tokenParsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return &crypt.ParsePrivateKey().PublicKey, nil
	})
	if err != nil {
		return false, err
	}

	if tokenParsed.Valid {
		return true, nil
	}

	return true, nil
}
