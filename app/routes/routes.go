package routes

import (
  "app/contorollers/todos"
  "app/contorollers/user"
  "app/contorollers/utils"
  "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  api := app.Group("api")
  admin := api.Group("admin")
  todo := api.Group("todo")

  // ====================
  // == Admin End Point

  // -- Get
  admin.Get("me", user.GetUserData)

  // -- Post
  admin.Post("login", user.Login)
  admin.Post("register", user.Register)
  admin.Post("verify-user", user.VerifyUser)
  admin.Post("email-existence", utils.ValidateEmailExistence)

  // -- Delete
  admin.Delete("user-delete", user.DeleteUser)
  admin.Delete("cleanup-all-user", user.CleanUpUserDatabase)

  // ====================
  // == ToDo End Point

  // -- Get
  todo.Get("/", todos.GetToDo)
  todo.Get("/all", todos.GetAllToDo)

  // -- Post
  todo.Post("create", todos.AddToDo)

  // ====================
  // == Index page

  // Get
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
  })

}
