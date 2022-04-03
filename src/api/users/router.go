package users

import (
	"discord-clone/src/api/users/create"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App) {
	newRouter := fiber.New()

	newRouter.Post("/", create.CreateUserController)

	router.Mount("/users", newRouter)
}
