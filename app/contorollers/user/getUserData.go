package user

import (
  "app/contorollers/utils"
  "app/database"
  "app/ent/user"
  "github.com/gofiber/fiber/v2"
  "strconv"
)

func GetUserData(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Request body parsing error.",
        "body":    err,
      },
      "user": nil,
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
      "user": nil,
    })
  }

  userID, _ := strconv.Atoi(payload.Subject)
  user, err := database.Client.User.Query().Where(user.ID(userID)).Only(ctx.Context())

  if err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Couldn't get user data.",
        "body":    err,
      },
      "user": nil,
    })
  }

  responseUser := utils.GenerateUserStruct(user, jwtString)

  return ctx.JSON(fiber.Map{
    "user": responseUser,
  })
}
