package create

import (
	"discord-clone/src/config"
	"discord-clone/src/database/models"
	"discord-clone/src/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type CreateUserResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"deletedAt"`
}

func CreateUserService(ctx *fiber.Ctx, body CreateUserBody) (CreateUserResponse, utils.Error) {
	var possibleUser models.User
	config.DB.Where(models.User{Email: body.Email}).FirstOrInit(&possibleUser)
	if possibleUser.Id != "" {
		return CreateUserResponse{}, utils.Error{
			Status:  fiber.StatusBadRequest,
			Message: "this user already exist",
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	newUser := models.User{
		Id:       utils.GenerateUUID(),
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}
	config.DB.Create(&newUser)

	return CreateUserResponse{
		Id:        newUser.Id,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	}, utils.Error{}
}
