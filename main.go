package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/liridonrama/microservices-with-go-users/src/database"
	"github.com/liridonrama/microservices-with-go-users/src/routes"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	routes.Setup(app)

	app.Listen(":8000")
}
