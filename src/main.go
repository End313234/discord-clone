package main

import (
	"discord-clone/src/database"
	"discord-clone/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func hello(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Hello World!",
	})
}

func main() {
	app := fiber.New()
	database.Connect()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
