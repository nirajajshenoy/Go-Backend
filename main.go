package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nirajajshenoy/go_backend/controllers"
)

func main() {
	fmt.Println("Welcome to the introduction")
	app := fiber.New()
	app.Use((logger.New()))
	fmt.Println("Calling Controller function:")
	// controllers.Controller() // Call the function from controllers package
	app.Get("/hello", controllers.GetController)
	app.Post("/hello", controllers.PostController)
	app.Delete("/hello", controllers.DeleteController)



	app.Get("/", func(o *fiber.Ctx) error {
		return o.SendString("Get All Data")
	})

	app.Post("/", func(o *fiber.Ctx) error {
		return o.SendString("To post a data")
	})

	app.Get("/api", func(o *fiber.Ctx) error {
		return o.SendString(" Get particular Data ")
	})

	app.Delete("/api", func(o *fiber.Ctx) error {
		return o.SendString("Delete a data")
	})

	app.Listen(":8080")
}
