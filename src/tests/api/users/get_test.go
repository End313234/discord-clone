package users_test

import (
	"discord-clone/src/config"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	chk := assert.New(t)

	response, err := http.Get(config.BASE_URL + "/users/" + config.EXAMPLE_USER_ID)
	chk.NoError(err)
	chk.Equal(fiber.StatusOK, response.StatusCode)
}

func TestGetUser_Failure_NotFound(t *testing.T) {
	chk := assert.New(t)

	response, err := http.Get(config.BASE_URL + "/users/random-id")
	chk.NoError(err)
	chk.Equal(fiber.StatusNotFound, response.StatusCode)
}
