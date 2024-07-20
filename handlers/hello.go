package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	greeting := os.Getenv("GREETING")
	return c.SendString(greeting + ", World!")
}
