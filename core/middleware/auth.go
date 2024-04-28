package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/auth"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get the JWT token from the Authorization header
	tokenString := c.Get("Authorization")
	// Check if the token is missing
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "missing token"})
	}
	// Check if the token starts with 'Bearer '
	if strings.HasPrefix(tokenString, "Bearer ") {
		// Remove 'Bearer ' from the token string
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		// Return an error if the token does not start with 'Bearer '
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "token string should not contain 'Bearer '"})
	}

	// Decode the JWT token
	claims, err := auth.DecodeJwtToken(tokenString)
	if err != nil {
		// Return an error if failed to decode the token
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// Set user claims to the context for further use
	c.Locals("user", claims)
	// Proceed to the next middleware or route handler
	return c.Next()
}
