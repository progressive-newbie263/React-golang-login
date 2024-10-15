package routes

import (
  "github.com/gofiber/fiber/v2"
	"server/controllers"
)

func SetUp(app *fiber.App) {
	//app.Get("/api/user", controllers.User)
	app.Post("/api/register", controllers.Register) 	//POST method
	app.Post("/api/login", controllers.Login)					//POST method
	app.Get("/api/user", controllers.User)					//GET method
	app.Post("/api/logout", controllers.Logout)					//POST method
}