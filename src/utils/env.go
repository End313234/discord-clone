package utils

import "github.com/joho/godotenv"

func GetEnv(key string) string {
	env, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}

	data, _ := env[key]
	return data
}
