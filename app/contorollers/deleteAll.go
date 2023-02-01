package contorollers

import (
  "app/database"
  "github.com/gofiber/fiber/v2"
)

func CleanUpUserDatabase(ctx *fiber.Ctx) error {
  if _, err := database.Client.User.Delete().Exec(ctx.UserContext()); err != nil {
    return ctx.Status(503).JSON(fiber.Map{
      "message": "Failed to delete all users",
      "err":     err,
    })
  }

  return ctx.Status(200).JSON(fiber.Map{
    "message": "Deleted all users.",
  })
}
