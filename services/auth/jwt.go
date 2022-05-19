package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

const CookieName = "token"

var jwtKey = []byte("secret_key")

func generateToken(email string, expiration time.Time) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GetAuthCookie(email string) *http.Cookie {
	expiration := time.Now().Add(12 * time.Hour)
	token, _ := generateToken(email, expiration)
	return &http.Cookie{
		Name:    CookieName,
		Value:   token,
		Expires: expiration,
	}
}

func Authenticate(c *http.Cookie) (string, error) {
	token := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		return "", errors.New("unauthorised")
	}
	return claims.Email, nil
}
