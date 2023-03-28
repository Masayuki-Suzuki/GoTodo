package todos

import (
  "app/database"
  "github.com/gofiber/fiber/v2"
  "github.com/joho/godotenv"
)

func DeleteAllToDo(ctx *fiber.Ctx) error {
  env, _ := godotenv.Read()
  var data map[string]string

  if data["secret_param"] == env["ADMIN_SECRET"] {
    if _, err := database.Client.Todo.Delete().Exec(ctx.Context()); err != nil {
      return ctx.Status(503).JSON(fiber.Map{
        "message": "Failed to delete all users",
        "err":     err,
      })
    }

    return ctx.Status(200).JSON(fiber.Map{
      "message": "Deleted all ToDos.",
    })
  } else {
    return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "message": "Unauthorised.",
    })
  }

  return ctx.Status(fiber.StatusTeapot).JSON(fiber.Map{
    "message": "I am tea pot ;)",
  })
}
