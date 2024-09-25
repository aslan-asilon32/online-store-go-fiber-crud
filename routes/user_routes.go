package routes

import (
	"online-store-go-fiber-crud/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}
