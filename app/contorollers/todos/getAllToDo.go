package todos

import (
  "app/contorollers/utils"
  "app/database"
  "app/ent/user"
  "fmt"
  "github.com/gofiber/fiber/v2"
  "strconv"

  todos "app/ent/todo"
)

func GetAllToDo(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Request body parsing error.",
        "body":    err,
      },
      "todos": nil,
    })
  }

  if data["token"] == "" {
    ctx.Status(fiber.StatusUnauthorized)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Token is required.",
      },
    })
  }

  jwtString := data["token"]
  payload, err := utils.ValidateToken(ctx, jwtString)

  if err != nil {
    ctx.Status(fiber.StatusUnauthorized)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Token is unauthorised",
        "body":    err,
      },
    })
  }

  userID, _ := strconv.Atoi(payload.Subject)
  todos, err := database.Client.Todo.Query().Where(todos.HasUserWith(user.ID(userID))).All(ctx.Context())

  if err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Database connection error.",
        "body":    err,
      },
      "todos": nil,
    })
  }

  // ToDo: May add the pagination.

  fmt.Println(todos)
  return ctx.JSON(fiber.Map{
    "todos": todos,
  })
}
