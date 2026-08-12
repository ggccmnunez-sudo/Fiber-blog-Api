package utils

import (
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(userId uint, username string) (string, error) {

	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	secret := os.Getenv("firstTrialJWT")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		secret := os.Getenv("firstTrialJWT")
		return []byte(secret), nil
	})
}
