package create

import (
	"discord-clone/src/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,gte=6" json:"password,omitempty"`
}

func CreateUserController(ctx *fiber.Ctx) error {
	parsedBody := &CreateUserBody{}
	err := utils.ValidateRequestBody(ctx.Body(), ctx, parsedBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.Error{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	newUser, httpError := CreateUserService(ctx, *parsedBody)
	if (httpError != utils.Error{}) {
		return ctx.Status(httpError.Status).JSON(httpError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(newUser)
}
