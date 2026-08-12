package middleware

import (
	"fiber-blog-api/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c fiber.Ctx) error {
		key := c.Get("Authorization")

		if key == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "missing authorization header",
			})
		}
		bearer := strings.Split(key, " ")

		if len(bearer) != 2 || bearer[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid authorization header",
			})
		}

		parsed, err := utils.ParseToken(bearer[1])
		if err != nil || !parsed.Valid {
			return c.Status(401).JSON(fiber.Map{
				"message": "invalid or expired token",
			})
		}
		claims := parsed.Claims.(jwt.MapClaims)
		c.Locals("user_id", uint(claims["user_id"].(float64)))

		return c.Next()
	}

}
