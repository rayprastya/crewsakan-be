package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/menus", controllers.PostMenu)
	app.Post("/users", controllers.CreateUser)
	app.Get("/menus/:userID", controllers.GetMenusByUserID)
}
