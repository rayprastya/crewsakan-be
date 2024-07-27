package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/db"
	"github.com/rayprastya/boost-daily/app/models"
)

func PostOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return err
	}
	result := db.GetDB().Create(&order)
	if result.Error != nil {
		log.Printf("Error creating order: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating order",
		})
	}
	return c.JSON(order)
}
