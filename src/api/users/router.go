package users

import (
	"discord-clone/src/api/users/create"
	"discord-clone/src/api/users/get"
	getall "discord-clone/src/api/users/get_all"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App) {
	newRouter := fiber.New()

	newRouter.Get("/", getall.GetAllUsersController)
	newRouter.Get("/:id", get.GetUserController)
	newRouter.Post("/", create.CreateUserController)

	router.Mount("/users", newRouter)
}
