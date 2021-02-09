package middlewares

import (
	"mypress/database"
	"mypress/models"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware middleware
func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("mp-token")

	if len(token) == 0 {
		return c.Redirect("/mp-admin/login")
	}

	user := models.User{
		Token: token,
	}
	result := database.Get().Where(user).Limit(1).Find(&user)
	if result.RowsAffected == 1 {
		c.Locals("user", user)
		return c.Next()
	}

	return c.Redirect("/mp-admin/login")
}
