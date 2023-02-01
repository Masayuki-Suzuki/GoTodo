package routes

import (
  "app/contorollers/user"
  "app/contorollers/utils"
  "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  api := app.Group("api")
  admin := api.Group("admin")

  //---------------------------
  // Admin End Point

  // Get
  //admin.Get("me")

  // Post
  admin.Post("login", user.Login)
  admin.Post("register", user.Register)
  admin.Post("email-existence", utils.ValidateEmailExistence)

  // Delete
  admin.Delete("user-delete", user.DeleteUser)
  admin.Delete("cleanup-all-user", user.CleanUpUserDatabase)

  // Get
  // Index page
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
  })

}
