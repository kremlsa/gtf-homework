package pages

import (
	"github.com/gofiber/fiber/v2"
)

type Category string

func PageHandler(router fiber.Router) {
	router.Get("/", homeApi)
}

func homeApi(c *fiber.Ctx) error {
	categories := []Category{"Еда", "Животные", "Машины", "Спорт", "Музыка", "Технологии", "Прочее"}
	data := struct {
		Categories []Category
	}{Categories: categories}
	return c.Render("home", data)
}
