package pages

import (
	"github.com/gofiber/fiber/v2"
)

func PageHandler(router fiber.Router) {
	router.Get("/", firstApi)
}

func firstApi(c *fiber.Ctx) error {

	return c.SendString("response")
}
