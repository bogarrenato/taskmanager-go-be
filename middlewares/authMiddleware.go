package middlewares

import (
	"fmt"
	"taskmanagerapp/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	fmt.Println("jön a cookie")
	fmt.Println(cookie)

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}
