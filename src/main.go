package main

import (
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
	app.Get("/", hello)
	log.Fatal(app.Listen(":3000"))
}
