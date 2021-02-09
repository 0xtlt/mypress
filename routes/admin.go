package routes

import "github.com/gofiber/fiber/v2"

// AdminIndex func
func AdminIndex(c *fiber.Ctx) error {
	return c.Send([]byte("ok"))
}

// AdminLogin func
func AdminLogin(c *fiber.Ctx) error {
	return c.Send([]byte("ok login"))
}
