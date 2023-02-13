package user

import (
  "app/contorollers/utils"
  "github.com/gofiber/fiber/v2"
)

func VerifyUser(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Request body parsing error.",
        "body":    err,
      },
      "is_authorised": false,
    })
  }

  jwtString := data["token"]
  _, err := utils.ValidateToken(ctx, jwtString)

  if err != nil {
    ctx.Status(fiber.StatusUnauthorized)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Token is unauthorised.",
        "body":    err,
      },
      "is_authorised": false,
    })
  }

  return ctx.JSON(fiber.Map{
    "error":         nil,
    "is_authorised": true,
  })
}
