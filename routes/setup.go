package routes

import (
	"mypress/controllers"
	"mypress/database"
	"mypress/models"
	"mypress/utils"

	"github.com/gofiber/fiber/v2"
)

// Setup route
func Setup(c *fiber.Ctx) error {
	if controllers.CheckWebsite() {
		return c.Redirect("/")
	}

	return c.Render("setup", fiber.Map{
		"url": ".",
	})
}

// SetupPost route
func SetupPost(c *fiber.Ctx) error {
	if controllers.CheckWebsite() {
		return c.Redirect("/")
	}

	form, err := c.MultipartForm()

	if err == nil {
		title := form.Value["title"]
		if len(title) == 0 {
			return c.Redirect("/setup")
		}

		description := form.Value["description"]
		if len(description) == 0 {
			return c.Redirect("/setup")
		}

		userDescription := form.Value["udescription"]
		if len(userDescription) == 0 {
			return c.Redirect("/setup")
		}

		url := form.Value["url"]
		if len(url) == 0 {
			return c.Redirect("/setup")
		}

		logos := form.File["logo"]
		if len(logos) == 0 || !utils.Includes(logos[0].Header["Content-Type"][0], []string{"image/png", "image/jpeg", "image/jpg"}) {
			return c.Redirect("/setup")
		}
		logo := logos[0]
		logoID := controllers.CreateMediaPicture(logo)

		banners := form.File["banner"]
		if len(banners) == 0 || !utils.Includes(banners[0].Header["Content-Type"][0], []string{"image/png", "image/jpeg", "image/jpg"}) {
			return c.Redirect("/setup")
		}
		banner := banners[0]
		bannerID := controllers.CreateMediaPicture(banner)

		font := form.Value["font"]
		if len(font) == 0 {
			return c.Redirect("/setup")
		}

		username := form.Value["username"]
		if len(username) == 0 {
			return c.Redirect("/setup")
		}

		email := form.Value["email"]
		if len(email) == 0 {
			return c.Redirect("/setup")
		}

		password := form.Value["password"]
		if len(password) == 0 {
			return c.Redirect("/setup")
		}

		avatars := form.File["banner"]
		if len(avatars) == 0 || !utils.Includes(avatars[0].Header["Content-Type"][0], []string{"image/png", "image/jpeg", "image/jpg"}) {
			return c.Redirect("/setup")
		}
		avatar := avatars[0]
		avatarID := controllers.CreateMediaPicture(avatar)

		theme := models.Theme{
			Name: "To The Moon!",
		}
		database.Get().Create(&theme)

		website := models.Website{
			Title:       title[0],
			Description: description[0],
			URL:         url[0],
			Font:        font[0],
			LogoID:      int(logoID),
			BannerID:    int(bannerID),
			ThemeID:     int(theme.ID),
		}

		database.Get().Create(&website)

		user := models.User{
			Username:    username[0],
			Email:       email[0],
			Password:    password[0],
			Token:       "",
			Description: description[0],
			IsAdmin:     true,
			AvatarID:    int(avatarID),
		}
		controllers.CreateUser(user)
	}

	return c.Redirect("/")
}
