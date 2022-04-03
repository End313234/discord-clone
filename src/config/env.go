package config

import "discord-clone/src/utils"

var BASE_URL, EXAMPLE_USER_ID, PORT string

func InitEnv(filename string) error {
	BASE_URL = utils.GetEnv("BASE_URL", filename)
	EXAMPLE_USER_ID = utils.GetEnv("EXAMPLE_USER_ID", filename)
	PORT = utils.GetEnv("PORT", filename)
	return nil
}
