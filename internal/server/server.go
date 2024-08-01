package server

import (
	"net/http"

	"movie-matcher/internal/config"
	"movie-matcher/internal/server/handlers"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/storage"
	"movie-matcher/internal/utilities"
	"movie-matcher/internal/views/not_found"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Setup(settings config.Settings) *fiber.App {
	app := createFiberApp()
	setupMiddleware(app)
	setupHealthCheck(app)
	service := createService(settings)
	setupCaching(app)
	setupRoutes(app, service)
	return app
}

func createFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder:       go_json.Marshal,
		JSONDecoder:       go_json.Unmarshal,
		ErrorHandler:      utilities.ErrorHandler,
		PassLocalsToViews: true,
	})
}

func setupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}

func createService(settings config.Settings) *handlers.Service {
	return handlers.NewService(
		storage.NewPostgresDB(settings.Database),
		omdb.NewCachedClient(),
	)
}

func setupRoutes(app *fiber.App, service *handlers.Service) {
	app.Route("/", func(r fiber.Router) {
		r.Get("favicon.ico", x404)
		r.Get("/deps/flowbite.min.js.map", x404)
		r.Get("", service.Home)
		r.Post("register", service.Register)
		r.Get("token", service.Token)
		r.Get("chart", service.Chart)
		r.Get("status", service.Status)
		r.Route(":token", func(r fiber.Router) {
			r.Get("prompt", service.Prompt)
			r.Post("submit", service.Submit)
		})
		r.Route("frontend", func(r fiber.Router) {
			r.Get("movies", service.Frontend)
		})
	})
}

func x404(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotFound)
}
