package contorollers

import (
  "app/database"
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
)

type UserDataStruct struct {
  id           int
  firstName    string
  lastName     string
  fullName     string
  emailAddress string
}

func Register(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    return err
  }

  pwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 16)

  result, _ := database.Client.User.
    Create().
    SetFirstName(data["firstName"]).
    SetLastName(data["lastName"]).
    SetEmail(data["emailAddress"]).
    SetPassword(string(pwd)).
    Save(ctx.Context())

  var userData UserDataStruct

  userData = UserDataStruct{
    id:           result.ID,
    firstName:    result.FirstName,
    lastName:     result.LastName,
    fullName:     result.FirstName + " " + result.LastName,
    emailAddress: result.Email,
  }

  return ctx.Status(200).JSON(fiber.Map{
    "data": userData,
  })
}
