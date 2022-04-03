package users_test

import (
	"discord-clone/src/config"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	chk := assert.New(t)

	response, err := http.Get(config.BASE_URL + "/users")
	chk.NoError(err)
	chk.Equal(fiber.StatusOK, response.StatusCode)
}
