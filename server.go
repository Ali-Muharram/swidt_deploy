package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Enable Prefork
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋! From Prefork mode.")
	})

	log.Fatal(app.Listen(":3000"))
}
