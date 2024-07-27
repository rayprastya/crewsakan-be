package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	openai "github.com/sashabaranov/go-openai"
)

func GetSummary(c *fiber.Ctx) error {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	var jsonData map[string]interface{}
	if err := c.BodyParser(&jsonData); err != nil {
		return err
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: jsonData["content"].(string),
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("Chat Completion error: %v \n\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": http.StatusInternalServerError,
			"summary":    "Something went wrong...",
		})
	}

	return c.JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"summary":    resp.Choices[0].Message.Content,
	})
}
