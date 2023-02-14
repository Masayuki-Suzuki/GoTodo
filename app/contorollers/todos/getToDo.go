package todos

import (
  "app/contorollers/utils"
  "app/database"
  "app/ent/todo"
  "github.com/gofiber/fiber/v2"
  "strconv"
)

func GetToDo(ctx *fiber.Ctx) error {
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

  if data["token"] == "" {
    ctx.Status(fiber.StatusUnauthorized)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Token is required.",
      },
    })
  }

  jwtString := data["token"]
  _, err := utils.ValidateToken(ctx, jwtString)

  if err != nil {
    ctx.Status(fiber.StatusUnauthorized)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Token is unauthorised",
        "body":    err,
      },
    })
  }

  if data["todoID"] == "" {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "ToDo ID is required.",
      },
    })
  }

  todoID, _ := strconv.Atoi(data["todoID"])
  todoItem, err := database.Client.Todo.Query().Where(todo.ID(todoID)).Only(ctx.Context())

  if err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Cannot find Todo",
        "todo_id": todoID,
        "body":    err,
      },
    })
  }

  return ctx.JSON(fiber.Map{
    "todo": todoItem,
  })

}
