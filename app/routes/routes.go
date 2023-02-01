package routes

import (
  "app/contorollers"
  "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  api := app.Group("api")
  admin := api.Group("admin")

  //---------------------------
  // Admin End Point

  // Post
  admin.Post("login", contorollers.Login)
  admin.Post("register", contorollers.Register)
  admin.Post("email-existence", contorollers.ValidateEmailExistence)

  // Delete
  admin.Delete("user-delete", contorollers.DeleteUser)
  admin.Delete("cleanup-all-user", contorollers.CleanUpUserDatabase)

  // Get
  // Index page
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
  })

}
