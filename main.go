package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-redis/repository"
)

func main() {
	// reids
	repository.SetupRedis()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8080")
}