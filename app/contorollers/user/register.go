package user

import (
  "app/contorollers/utils"
  "app/database"
  "app/types"
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
)

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

  token, err := utils.GetToken(result.ID)
  if err != nil {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "emailAddress": data["emailAddress"],
        "message":      "Email or Password is incorrect.",
      },
    })
  }

  var userData types.UserDataStruct

  userData = types.UserDataStruct{
    Id:           result.ID,
    FirstName:    result.FirstName,
    LastName:     result.LastName,
    FullName:     result.FirstName + " " + result.LastName,
    EmailAddress: result.Email,
    Token:        token,
  }

  return ctx.Status(200).JSON(fiber.Map{
    "data": userData,
  })
}
