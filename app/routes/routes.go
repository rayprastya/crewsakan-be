package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/menus", controllers.PostMenu)
	app.Get("/menus/:merchant_id", controllers.GetMenusByMerchantID)
	app.Post("/merchants", controllers.CreateMerchant)
	app.Get("/merchants/:id", controllers.GetMerchantByID)
	app.Post("/login", controllers.LoginMerchant)
	app.Post("/orders", controllers.PostOrder)
	app.Post("/wishlists", controllers.PostWishlist)
	app.Get("/wishlists/:merchant_id", controllers.GetWishlistByMerchantID)
}
