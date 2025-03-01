package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kremlsa/gtf-homework/config"
	"github.com/kremlsa/gtf-homework/internal/pages"
)

func main() {
	app := fiber.New()
	appConf := config.NewAppConfig()
	pages.PageHandler(app)

	app.Listen(fmt.Sprintf(":%d", appConf.Port))
}
