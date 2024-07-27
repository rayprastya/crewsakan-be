package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/db"
	"github.com/rayprastya/boost-daily/app/models"
)

func PostWishlist(c *fiber.Ctx) error {
	var wishlist models.Wishlist
	if err := c.BodyParser(&wishlist); err != nil {
		return err
	}
	// log.Printf("wishlist: %v", wishlist)
	result := db.GetDB().Create(&wishlist)
	if result.Error != nil {
		log.Printf("Error creating wishlist: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating wishlist",
		})
	}
	return c.JSON(wishlist)
}

func GetWishlistByMerchantID(c *fiber.Ctx) error {
	merchantID := c.Params("merchant_id")
	var wishlists []models.Wishlist
	result := db.GetDB().Where("merchant_id = ?", merchantID).Find(&wishlists)
	if result.Error != nil {
		log.Printf("Error fetching wishlists: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching wishlists",
		})
	}
	return c.JSON(wishlists)
}
