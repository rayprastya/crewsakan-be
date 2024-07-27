package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/menu", controllers.PostMenu)
	app.Post("/user", controllers.CreateUser)
	app.Get("/menu/:userID", controllers.GetMenusByUserID)
}
