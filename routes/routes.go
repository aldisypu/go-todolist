package routes

import (
	"github.com/aldisypu/go-todolist/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Post("/api/todos", controller.Create)
	app.Get("/api/todos", controller.FindAll)
	app.Get("/api/todos/:todoId", controller.FindById)
	app.Put("/api/todos/:todoId", controller.Update)
	app.Delete("/api/todos/:todoId", controller.Delete)
}
