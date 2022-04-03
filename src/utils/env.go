package utils

import (
	"github.com/joho/godotenv"
)

func GetEnv(key string, filename string) string {
	env, err := godotenv.Read(filename)
	if err != nil {
		panic(err)
	}

	data, _ := env[key]
	return data
}
