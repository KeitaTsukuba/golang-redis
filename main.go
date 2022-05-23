package main

import (
	"golang-redis/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// reids
	repository.SetupRedis()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("users/:uuid", getUserList)

	app.Listen(":8080")
}

func getUserList(c *fiber.Ctx) error {
	// リクエストからIDを取得
	uuid := c.Params("uuid")

	// redisからデータを取得
	userList, err := repository.GetUserList(uuid)

	if err != nil {
		panic(err)
	}

	return c.JSON(userList)
}