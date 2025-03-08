package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func PageHandler(router fiber.Router, logger *slog.Logger) {
	router.Get("/", firstApi)
}

func firstApi(c *fiber.Ctx) error {

	return c.SendString("response")
}
