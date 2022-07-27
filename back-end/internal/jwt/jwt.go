package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Key     string
	Timeout time.Duration
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(user string, pass string) (string, error) {
	claims := &Claims{
		Username: user,
		Password: pass,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.Timeout * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.Key)
}

func (j *JWT) ParseToken(jwtToken string) (bool, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})

	return tkn.Valid, err
}
