package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/boost-daily/app/db"
	"github.com/rayprastya/boost-daily/app/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	result := db.GetDB().Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating user",
		})
	}
	return c.JSON(user)
}
