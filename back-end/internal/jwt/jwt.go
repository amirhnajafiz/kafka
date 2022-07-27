package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(user string, pass string, key string, timeout time.Duration) (string, error) {
	claims := &JWT{
		Username: user,
		Password: pass,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeout * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func ParseToken(jwtToken string, key string) (bool, error) {
	claims := &JWT{}
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	return tkn.Valid, err
}
