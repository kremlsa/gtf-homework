package pages

import "github.com/gofiber/fiber/v2"

type PageHandler struct {
	router fiber.Router
}

func FirstHandler(router fiber.Router) {
	h := &PageHandler{
		router: router,
	}
	h.router.Get("/", h.firstApi)
}

func (p *PageHandler) firstApi(c *fiber.Ctx) error {
	return c.SendString("response")
}
