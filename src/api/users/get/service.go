package get

import (
	"discord-clone/src/config"
	"discord-clone/src/database/models"
	"discord-clone/src/utils"

	"github.com/gofiber/fiber/v2"
)

type GetUserResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"deletedAt"`
	Password  string `json:"-"`
}

func GetUserService(ctx *fiber.Ctx) (GetUserResponse, utils.Error) {
	var user models.User

	userId := ctx.Params("id")
	if userId == "" {
		return GetUserResponse{}, utils.Error{
			Status:  fiber.StatusBadRequest,
			Message: "missing id param",
		}
	}

	config.DB.Where(models.User{Id: userId}).First(&user)
	if user.Id == "" {
		return GetUserResponse{}, utils.Error{
			Status:  fiber.StatusNotFound,
			Message: "user not found",
		}
	}

	return GetUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, utils.Error{}
}
