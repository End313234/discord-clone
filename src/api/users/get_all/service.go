package getall

import (
	"discord-clone/src/config"
	"discord-clone/src/database/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"deletedAt"`
	Password  string `json:"-"`
}

type GetAllUsersResponse = []User

func GetAllUsersService(ctx *fiber.Ctx) GetAllUsersResponse {
	var users []models.User
	config.DB.Where(models.User{}).First(&users)

	var filteredUsers GetAllUsersResponse
	usersInBytes, _ := json.Marshal(users)
	json.Unmarshal(usersInBytes, &filteredUsers)

	return filteredUsers
}
