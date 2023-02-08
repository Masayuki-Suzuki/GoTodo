package user

import (
  "app/contorollers/utils"
  "app/database"
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Request body parsing error.",
        "body":    err,
      },
    })
  }

  pwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 16)

  result, _ := database.Client.User.
    Create().
    SetFirstName(data["firstName"]).
    SetLastName(data["lastName"]).
    SetEmail(data["emailAddress"]).
    SetPassword(string(pwd)).
    Save(ctx.Context())

  token, err := utils.GetToken(result.ID, result.Email)
  if err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "The system get an internal server error. Please try again later.",
      },
    })
  }

  userData := utils.GenerateUserStruct(result, token)

  return ctx.Status(200).JSON(fiber.Map{
    "user": userData,
  })
}
