package todos

import (
  "app/contorollers/utils"
  "app/database"
  "github.com/gofiber/fiber/v2"
  "strconv"
  "time"
)

func AddToDo(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Request body parsing error.",
        "body":    err,
      },
      "todo": nil,
    })
  }

  if data["title"] == "" {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Title is required.",
      },
      "todo": nil,
    })
  }

  jwtString := data["token"]
  payload, err := utils.ValidateToken(ctx, jwtString)

  if err != nil {
    ctx.Status(fiber.StatusUnauthorized)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Token is unauthorised.",
        "body":    err,
      },
      "todo": nil,
    })
  }

  userID, _ := strconv.Atoi(payload.Subject)
  createdAt := time.Now().UTC().String()

  todo, err := database.Client.Todo.
    Create().
    SetTitle(data["title"]).
    SetDescription(data["description"]).
    SetStatus("ToDo").
    SetDueDate(data["dueDate"]).
    SetCreatedAt(createdAt).
    SetCompletedAt("").
    SetUserID(userID).
    Save(ctx.Context())

  if err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Database connection error.",
        "body":    err,
      },
      "todo": nil,
    })
  }

  return ctx.JSON(fiber.Map{
    "todo": todo,
  })
}
