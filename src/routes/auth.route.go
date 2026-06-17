package routes

import (
	"todo/src/controllers"
	"todo/src/middleware"

	"github.com/gofiber/fiber/v3"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser) //auth/register
	auth.Post("/login", controllers.LoginUser) //auth/login
	auth.Post("/logout", middleware.AuthMiddleware, controllers.LogoutUser) //auth/logout
}