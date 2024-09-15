package routes

import (
	"url_shortener/api/database"

	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	// get the short URL
	url := c.Params("url")

	// query the db to find the original URL
	// if a match is found increment the redirect counter and redirect the original URL
	r := database.CreateClient(0)
}
