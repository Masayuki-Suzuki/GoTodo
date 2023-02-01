package contorollers

import (
  "app/database"
  "app/ent/user"
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    return ctx.Status(422).JSON(fiber.Map{
      "error": err,
    })
  }

  fmt.Println(data["emailAddress"])

  //result, err := database.Client.
  //  User.
  //  Query().
  //  Where(user.Email(data["emailAddress"])).
  //  All(ctx.Context())

  user, err := database.Client.
    User.
    Query().
    Where(user.Email(data["emailAddress"])).
    Only(ctx.Context())

  fmt.Printf("failed querying user: %v\n", err)

  return ctx.Status(200).JSON(fiber.Map{
    "user": user,
  })
}
