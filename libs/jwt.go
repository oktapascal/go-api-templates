package libs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("APP_NAME"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		Email: email,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SIGNATURE_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
