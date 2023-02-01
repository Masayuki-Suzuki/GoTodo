package contorollers

import (
  "app/database"
  "app/ent/user"
  "github.com/gofiber/fiber/v2"
)

func ValidateEmailExistence(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    return err
  }

  result, err := database.Client.
    User.
    Query().
    Where(user.Email(data["emailAddress"])).
    All(ctx.Context())

  res := fiber.Map{
    "is_used": false,
    "error":   err,
  }

  if len(result) > 0 {
    res["is_used"] = true
  }

  return ctx.JSON(res)
}
