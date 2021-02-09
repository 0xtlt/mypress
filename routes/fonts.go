package routes

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Fonts route
func Fonts(c *fiber.Ctx) error {
	data, err := ioutil.ReadFile("./assets/fonts.css")
	if err != nil {
		log.Fatal("File reading error", err)
	}

	content := string(data)
	ns := strings.Replace(content, "{{URL}}", "", -1)

	c.Response().Header.SetContentType("text/css")

	return c.Send([]byte(ns))
}
