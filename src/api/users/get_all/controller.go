package getall

import "github.com/gofiber/fiber/v2"

func GetAllUsersController(ctx *fiber.Ctx) error {
	users := GetAllUsersService(ctx)
	return ctx.Status(fiber.StatusOK).JSON(users)
}
