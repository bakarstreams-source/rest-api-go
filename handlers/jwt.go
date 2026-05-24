package handlers

import (
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string) (string, error) {
    secret := os.Getenv("JWT_SECRET")

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString([]byte(secret))
}