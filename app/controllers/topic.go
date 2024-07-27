package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/db"
	"github.com/rayprastya/boost-daily/app/models"
)

func GetTopics(c *fiber.Ctx) error {
	var records []models.Topic

	result := db.GetDB().Find(&records)
	if result.Error != nil {
		log.Printf("Error fetching topics: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching topics",
		})
	}
	return c.JSON(records)
}
