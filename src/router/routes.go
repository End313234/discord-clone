package router

import (
	"discord-clone/src/api/users"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	router := fiber.New()
	router.Post("/users", users.Create)

	app.Mount("/", router)
}
