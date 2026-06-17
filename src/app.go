package src

import (
	"log"
	"todo/src/db"
	"todo/src/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()

	app.Get("/",func (c fiber.Ctx) error {
		return c.SendString("Welcome To Todo made with fiber")
	})

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	return app
}