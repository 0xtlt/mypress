package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CreateMedia route
func CreateMedia(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err == nil {
		// => *multipart.Form
		// Get all files from "documents" key:
		files := form.File["documents"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	return err
}
