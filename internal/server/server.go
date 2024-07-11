package server

import (
	"net/http"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/config"
	"movie-matcher/internal/server/handlers"
	"movie-matcher/internal/storage"
	"movie-matcher/internal/utilities"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Setup(settings config.Settings) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: utilities.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	service := handlers.NewService(
		storage.NewPostgresDB(settings.Database),
		&algo.MoviePrompter{},
	)

	app.Route("/",
		func(r fiber.Router) {
			r.Get("health", func(c *fiber.Ctx) error {
				return c.SendStatus(http.StatusOK)
			})
			r.Post("register", service.Register)
			r.Route(":nuid", func(r fiber.Router) {
				r.Get("token", service.Token)
			})
			r.Route(":token", func(r fiber.Router) {
				r.Get("prompt", service.Prompt)
				r.Post("submit", service.Submit)
			})
		},
	)

	return app
}
