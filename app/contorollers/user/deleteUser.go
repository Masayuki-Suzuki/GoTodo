package user

import (
  "app/database"
  "github.com/gofiber/fiber/v2"
  "reflect"
)

func DeleteUser(ctx *fiber.Ctx) error {
  data := make(map[string]int)

  if err := ctx.BodyParser(&data); err != nil {
    return ctx.Status(503).JSON(fiber.Map{
      "error": err,
    })
  }

  typeOfId := reflect.TypeOf(data["id"]).String()

  if typeOfId == "string" {
    return ctx.Status(503).JSON(fiber.Map{
      "message": "User ID must be a number.",
      "error":   "Server received User ID as string.",
    })
  }

  if data["id"] > 0 {
    if err := database.Client.User.DeleteOneID(data["id"]).Exec(ctx.Context()); err != nil {
      return ctx.Status(503).JSON(fiber.Map{
        "message": "Database cannot delete user",
        "user_id": data["id"],
        "error":   err,
      })
    }

    return ctx.Status(200).JSON(fiber.Map{
      "message": "Success deleting user.",
      "user_id": data["id"],
    })
  }

  return ctx.Status(503).JSON(fiber.Map{
    "message": "User ID is required.",
  })
}
