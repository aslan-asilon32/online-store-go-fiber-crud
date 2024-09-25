package main

import (
	"online-store-go-fiber-crud/config"
	"online-store-go-fiber-crud/middlewares"
	"online-store-go-fiber-crud/models"
	"online-store-go-fiber-crud/routes"

	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Connect to the database
	config.Connect()

	// Migrate the schema
	config.DB.AutoMigrate(&models.User{})

	// Inisialisasi aplikasi Fiber
	app := fiber.New(fiber.Config{

		Views: html.New(filepath.Join(".", "views"), ".html"),
	})

	// Middleware
	app.Use(middlewares.Logger())

	// Routes
	routes.UserRoutes(app)

	// Start the server
	app.Listen(":4000")
}
