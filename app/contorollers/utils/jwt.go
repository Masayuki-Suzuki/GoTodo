package utils

import (
  "github.com/golang-jwt/jwt/v4"
  "github.com/joho/godotenv"
  _ "github.com/joho/godotenv/autoload"
  "strconv"
  "time"
)

func GetToken(userID int, email string) (string, error) {
  env, err := godotenv.Read()

  payload := jwt.RegisteredClaims{
    Issuer:    email,
    Subject:   strconv.Itoa(userID),
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
  }

  token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(env["JWT_SECRET"]))
  return token, err
}

func ValidateToken(token string, ) bool {
  return false
}
