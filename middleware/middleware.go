package middleware

import (
	"github.com.vinay-kumar-ps/blogbackend/util"
	"github.com/gofiber/fiber/v2"
)

// IsAuthenticated checks if the user is authenticated via JWT
func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}
