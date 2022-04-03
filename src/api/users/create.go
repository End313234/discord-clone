package users

import (
	"discord-clone/src/config"
	"discord-clone/src/database/models"
	"discord-clone/src/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,gte=6" json:"password,omitempty"`
}

type CreateUserResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"deletedAt"`
}

func Create(ctx *fiber.Ctx) error {
	parsedBody := &CreateUserBody{}
	err := utils.ValidateRequestBody(ctx.Body(), ctx, parsedBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.Error{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var possibleUser models.User
	config.DB.Where(models.User{Email: parsedBody.Email}).FirstOrInit(&possibleUser)
	if possibleUser.Id != "" {
		return ctx.Status(fiber.StatusForbidden).JSON(utils.Error{
			Status:  fiber.StatusForbidden,
			Message: "user already exist",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(parsedBody.Password), 8)
	newUser := models.User{
		Id:       utils.GenerateUUID(),
		Name:     parsedBody.Name,
		Email:    parsedBody.Email,
		Password: string(hashedPassword),
	}
	config.DB.Create(&newUser)

	return ctx.Status(fiber.StatusOK).JSON(CreateUserResponse{
		Id:        newUser.Id,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	})
}
