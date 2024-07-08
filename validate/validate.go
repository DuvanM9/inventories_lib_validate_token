package validate

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(currentToken string) error {

	key := []byte(os.Getenv("TOKEN_KEY"))

	token, err := jwt.Parse(currentToken, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrorInvalitToken
	}

	///castear token claims con maplains
	claims, okClaims := token.Claims.(jwt.MapClaims)

	if !okClaims {
		return ErrorInvalitClaims
	}

	//validar la fecha de expiracion del token
	exp := int64(claims["exp"].(float64))
	if exp < time.Now().Unix() {
		return ErrorInvalitToken
	}

	return nil

}
