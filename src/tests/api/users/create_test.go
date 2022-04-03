package users_test

import (
	"bytes"
	"discord-clone/src/api/users/create"
	"discord-clone/src/constraints"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	chk := assert.New(t)
	requestBody := create.CreateUserBody{
		Name:     "test",
		Email:    "test@domain.com",
		Password: "123456",
	}

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(requestBody)
	chk.NoError(err)

	client := http.Client{}
	req, err := http.NewRequest(
		http.MethodPost,
		constraints.BASE_URL+"/users",
		&buffer,
	)
	chk.NoError(err)

	response, err := client.Do(req)
	chk.NoError(err)
	chk.Equal(fiber.StatusCreated, response.StatusCode)
}

func TestCreateUser_Failure(t *testing.T) {
	chk := assert.New(t)
	requestBody := create.CreateUserBody{
		Name:     "test",
		Email:    "test@domain.com",
		Password: "12345",
	}

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(requestBody)
	chk.NoError(err)

	client := http.Client{}
	req, err := http.NewRequest(
		http.MethodPost,
		constraints.BASE_URL+"/users",
		&buffer,
	)
	chk.NoError(err)

	response, err := client.Do(req)
	chk.NoError(err)
	chk.Equal(fiber.StatusBadRequest, response.StatusCode)
}
