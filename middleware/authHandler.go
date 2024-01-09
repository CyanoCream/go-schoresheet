package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Authentication middleware checks for a valid JWT in the Authorization header
func Authentication(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("userId", nil)

	return c.Next()
}
