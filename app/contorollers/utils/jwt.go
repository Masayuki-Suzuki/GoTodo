package utils

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
  "github.com/golang-jwt/jwt/v4"
  "github.com/joho/godotenv"
  _ "github.com/joho/godotenv/autoload"
  "strconv"
  "time"
)

func GetToken(userID int, email string) (string, error) {
  env, err := godotenv.Read()

  fmt.Println(strconv.Itoa(userID))

  payload := jwt.RegisteredClaims{
    Issuer:    email,
    Subject:   strconv.Itoa(userID),
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
  }

  token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(env["JWT_SECRET"]))
  return token, err
}

func ValidateToken(ctx *fiber.Ctx, jwtString string, ) (*jwt.RegisteredClaims, error) {
  env, _ := godotenv.Read()
  token, err := jwt.ParseWithClaims(jwtString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(env["JWT_SECRET"]), nil
  })

  if err != nil {
    return nil, err
  }

  payload := token.Claims.(*jwt.RegisteredClaims)
  return payload, nil
}
