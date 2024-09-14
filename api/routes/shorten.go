package routes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	URL               string        `json:"url"`
	CustomShortExpiry string        `json:"short`
	Expiry            time.Duration `json:expiry"`
	XRateRemaining    int           `json:"rate_limit"`
	XRateLimitRest    time.Duration `json:"rate_limit_reset"`
}

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if error := c.BodyParser(&body); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	//implement rate limiting
	// everytime a user queries, check if the IP is already in database,
	// if yes, decrement the calls remaining by one, else add the IP to database
	r2 := database

	// check if the input is an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	//check for the domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "haha... nice try",
		})
	}

	//enforce http, SSL

	body.URL = helpers.EnforceHTTP(body.URL)
}
