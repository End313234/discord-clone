package router

import (
	"discord-clone/src/api/users"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	router := fiber.New()

	users.RegisterRoutes(router)

	app.Mount("/", router)
}
