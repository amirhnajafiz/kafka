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
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(timeout * time.Minute).Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
