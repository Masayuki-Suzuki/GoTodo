package user

import (
  "app/contorollers/utils"
  "app/database"
  "app/ent"
  "app/ent/user"
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    return ctx.Status(422).JSON(fiber.Map{
      "error": err,
    })
  }

  user, err := database.Client.
    User.
    Query().
    Where(user.Email(data["emailAddress"])).
    Only(ctx.Context())

  if ent.IsNotFound(err) {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message":      "Cannot find the user.",
        "emailAddress": data["emailAddress"],
      },
    })
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
  if err != nil {
    ctx.Status(fiber.StatusBadRequest)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "emailAddress": data["emailAddress"],
        "message":      "Email or Password is incorrect.",
      },
    })
  }

  token, err := utils.GetToken(user.ID, user.Email)
  if err != nil {
    ctx.Status(fiber.StatusInternalServerError)
    return ctx.JSON(fiber.Map{
      "error": fiber.Map{
        "message": "Cannot get token",
      },
    })
  }

  responseUser := utils.GenerateUserStruct(user, token)

  return ctx.Status(200).JSON(fiber.Map{
    "user": responseUser,
  })
}
