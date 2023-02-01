package user

import "github.com/gofiber/fiber/v2"

func GetTheUser(ctx *fiber.Ctx) error {
  var data map[string]string

  if err := ctx.BodyParser(&data); err != nil {
    return ctx.Status(422).JSON(fiber.Map{
      "error": err,
    })
  }

  return nil
}
