package jwt

import (
	"errors"
	"fmt"
	"time"

	"api.default.indicoinnovation.pt/pkg/constants"
	"api.default.indicoinnovation.pt/pkg/crypt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtHeaders struct {
	Key   string
	Value string
}

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

var Headers = map[string]string{
	"alg": "RS512",
	"typ": "JWT",
}

func SetupClaims(userEmail string, customArgs ...JwtHeaders) jwt.MapClaims {
	accessTokenClaims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(constants.AccessTokenExpirationTime) * time.Minute).Unix(),
		"aud": constants.Audience,
		"jti": uuid.New().String(),
		"iss": userEmail,
		"sub": userEmail,
	}

	if len(customArgs) > 0 {
		fmt.Print("TODO")
	}

	return accessTokenClaims
}

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
