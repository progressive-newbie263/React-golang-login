package main

import (
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" //package for cors policy.
	_ "github.com/lib/pq"
)



func main() {
  psqlconn := "postgres://postgres:26032004@localhost/users-test?sslmode=disable"
  database.Connect(psqlconn)

  app := fiber.New()

  app.Use(cors.New(cors.Config{
    AllowCredentials: true,
    AllowOrigins: "http://localhost:8000, http://localhost:5173",
  }))

  routes.SetUp(app)

  app.Listen(":8000")
}