package get

import (
	"discord-clone/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUserController(ctx *fiber.Ctx) error {
	user, err := GetUserService(ctx)
	if (err != utils.Error{}) {
		return ctx.Status(err.Status).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
