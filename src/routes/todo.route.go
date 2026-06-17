package routes

import (
	"todo/src/controllers"
	"todo/src/middleware"

	"github.com/gofiber/fiber/v3"
)

func TodoRoutes(app *fiber.App) {
	todo := app.Group("/todo", middleware.AuthMiddleware)

	todo.Post("/", controllers.CreateTodo)
	todo.Get("/", controllers.GetTodos)
	todo.Delete("/:id", controllers.DeleteTodo)
	todo.Put("/:id", controllers.UpdateTodo)
}