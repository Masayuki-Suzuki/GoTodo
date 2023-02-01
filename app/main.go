package main

import (
  "app/database"
  "app/routes"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/template/html"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  database.Connect()
  defer database.Client.Close()

  // Initialize standard Go html template engine
  engine := html.New("./views", ".html")

  app := fiber.New(fiber.Config{
    Views: engine,
  })

  app.Use(cors.New(cors.Config{
    AllowCredentials: true,
  }))

  routes.Setup(app)
  log.Fatal(app.Listen(":3000"))
}
