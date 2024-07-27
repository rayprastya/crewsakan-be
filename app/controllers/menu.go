package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/db"
	"github.com/rayprastya/boost-daily/app/models"
)

func PostMenu(c *fiber.Ctx) error {
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return err
	}
	result := db.GetDB().Create(&menu)
	if result.Error != nil {
		log.Printf("Error creating menu: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating menu",
		})
	}
	return c.JSON(menu)
}

func GetMenusByUserID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	var menus []models.Menu
	result := db.GetDB().Where("user_id = ?", userID).Find(&menus)
	if result.Error != nil {
		log.Printf("Error fetching menus: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching menus",
		})
	}
	return c.JSON(menus)
}
