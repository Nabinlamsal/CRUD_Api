package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("123abc")

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token missing!",
			})
		}

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token!",
			})
		}

		// Token valid → move to next
		return c.Next()
	}
}
