package main

import (
	"log"
	"mypress/controllers"
	"mypress/database"
	"mypress/middlewares"
	"mypress/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	var websiteInstalled bool
	websiteInstalled = controllers.CheckWebsite()

	engine := html.New("./views", ".html")
	engine.Reload(true)
	engine.Debug(true)
	engine.Layout("embed")
	engine.Delims("{{", "}}")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.Get()

	app.Get("/_/fonts.css", routes.Fonts)
	app.Static("/_/", "./public")
	app.Get("/setup", routes.Setup)
	app.Post("/setup", routes.SetupPost)

	app.Use(func(c *fiber.Ctx) error {
		if websiteInstalled {
			return c.Next()
		}

		websiteInstalled = controllers.CheckWebsite()
		if websiteInstalled {
			return c.Next()
		}
		return c.Redirect("/setup")
	})

	admin := app.Group("/mp-admin")
	admin.Get("/login", routes.AdminLogin)
	admin.Use(middlewares.AuthMiddleware)
	admin.Get("/", routes.AdminIndex)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
