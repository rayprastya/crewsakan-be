package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/summary", controllers.GetSummary)
	app.Get("/topics", controllers.GetTopics)
}
