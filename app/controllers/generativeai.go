package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rayprastya/crewsakan-be/app/db"
	"github.com/rayprastya/crewsakan-be/app/models"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateText(client *openai.Client, content string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo-1106",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		req,
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func GenerativeAI(client *openai.Client, command string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var jsonData map[string]interface{}
		if err := c.BodyParser(&jsonData); err != nil {
			return err
		}

		var content string
		var respContent string
		var err error

		if command == "GetListMenu" {
			content = "I want you to recommend me a similar menu, " + jsonData["content"].(string) + "give me the list of the menu name only in json with recommendation as the only key"
			respContent, err = GenerateText(client, content)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Something went wrong...",
				})
			}

			type Recipe struct {
				Recommendation []string `json:"recommendation"`
			}

			var data Recipe
			if err := json.Unmarshal([]byte(respContent), &data); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Failed to parse response",
				})
			}

			return c.JSON(fiber.Map{
				"statusCode": http.StatusOK,
				"foodList":   data,
			})

		} else if command == "GetRecipe" {
			content = "how to make " + jsonData["content"].(string) + "? give me the list of all details steps and ingredients in json with ingredients and steps as the only key"
			respContent, err = GenerateText(client, content)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Something went wrong...",
				})
			}

			type Recipe struct {
				Ingredients []string `json:"ingredients"`
				Steps       []string `json:"steps"`
			}

			var data Recipe
			if err := json.Unmarshal([]byte(respContent), &data); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Failed to parse response",
				})
			}

			return c.JSON(fiber.Map{
				"statusCode": http.StatusOK,
				"summary":    data,
			})

		} else if command == "GetStep" {
			content = "how to make " + jsonData["content"].(string) + "give me the list of all the steps only in json"
			respContent, err = GenerateText(client, content)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Something went wrong...",
				})
			}

			type Steps struct {
				Steps []string `json:"steps"`
			}

			var data Steps
			if err := json.Unmarshal([]byte(respContent), &data); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Failed to parse response",
				})
			}

			return c.JSON(fiber.Map{
				"statusCode": http.StatusOK,
				"steps":      data,
			})

		} else if command == "RecommendNearby" {
			var jsonData map[string]interface{}
			if err := c.BodyParser(&jsonData); err != nil {
				return err
			}

			merchantID := jsonData["merchant_id"].(float64)
			fmt.Println("merchantID", merchantID)

			var menu_list []models.Menu
			if err := db.GetDB().
				Where("merchant_id = ?", merchantID).Find(&menu_list).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Merchant not found",
				})
			}

			fmt.Println("merchant", menu_list)

			// TODO: ambil jarak +- 5km dari merchant

			var orders []models.Order
			if err := db.GetDB().
				Preload("Menu").
				Where("order_datetime BETWEEN NOW() - INTERVAL '5 DAYS' AND NOW()").
				Find(&orders).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Error fetching orders",
				})
			}
			var merchantFood string

			for i := 0; i < len(menu_list); i++ {
				merchantFood += menu_list[i].Name
			}

			var userPreference string

			for i := 0; i < len(orders); i++ {
				userPreference += orders[i].Menu.Name + " preference: " + orders[i].Optional
			}

			content = "this is list of most buyed food in nearby market " + userPreference + " give me top 3 the list of recommended food based on that list but exclude the same recommendation as " + merchantFood + ". i want the response in json with recommendation as the only key"
			respContent, err = GenerateText(client, content)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Something went wrong...",
				})
			}

			type Recipe struct {
				Recommendation []string `json:"recommendation"`
			}

			var data Recipe
			if err := json.Unmarshal([]byte(respContent), &data); err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": http.StatusInternalServerError,
					"summary":    "Failed to parse response",
				})
			}

			fmt.Println("data", data)
			fmt.Println("respContent", respContent)

			return c.JSON(fiber.Map{
				"statusCode": http.StatusOK,
				"foodList":   data,
			})

		} else {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"statusCode": http.StatusBadRequest,
				"summary":    "Invalid command",
			})
		}
	}
}
