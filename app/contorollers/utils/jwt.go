package utils

import (
  "fmt"
  "github.com/golang-jwt/jwt/v4"
  "github.com/joho/godotenv"
  "strconv"
  "time"

  _ "github.com/joho/godotenv/autoload"
)

func GetToken(userID int) (string, error) {
  env, err := godotenv.Read()

  fmt.Println(env)

  payload := jwt.RegisteredClaims{
    Subject:   strconv.Itoa(userID),
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
  }
  token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
  return token, err
}
