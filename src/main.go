package main

import (
	"discord-clone/src/config"
	"discord-clone/src/database"
	"discord-clone/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.Connect()
	config.InitEnv(".env")

	router.SetupRoutes(app)
	log.Fatal(app.Listen(config.PORT))
}
