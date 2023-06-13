package jwt

import (
	"api.default.indicoinnovation.pt/clients/iam"
	"api.default.indicoinnovation.pt/pkg/crypt"
)

func Validate(token string) bool {
	client, context := iam.New()

	auth, err := client.ValidateJWT(context, token, crypt.ParsePrivateKeyToString())
	if err != nil {
		return false
	}

	return auth.Jwt != ""
}
