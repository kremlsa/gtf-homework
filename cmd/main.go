package main

import (
	"fmt"
	"os"

	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kremlsa/gtf-homework/configs"
	"github.com/kremlsa/gtf-homework/internal/pages"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {

	// Создаём логгер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// Добавляем конфигурацию приложения
	appConf := configs.NewAppConfig()

	// Инициализируем приложение
	app := fiber.New()

	// Подключаем логгер
	app.Use(slogfiber.New(logger))
	app.Use(recover.New())

	pages.PageHandler(app)

	app.Listen(fmt.Sprintf(":%d", appConf.Port))
}
