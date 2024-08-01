package server

import (
	"net/http"

	"movie-matcher/internal/config"
	"movie-matcher/internal/constants"
	"movie-matcher/internal/server/handlers"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/storage"
	"movie-matcher/internal/utilities"

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
	app.Get("/", service.Index)
	app.Get("/favicon.ico", x404)
	app.Get("/deps/flowbite.min.js.map", x404)
	app.Post("/register", service.Register)
	app.Get("/token", service.Token)
	app.Get("/chart", service.Chart)
	app.Get("/status", service.Status)

	app.Route("/challenges", func(r fiber.Router) {
		r.Get("/backend", service.Backend)
		r.Get("/frontend", func(c *fiber.Ctx) error {
			return c.Redirect(constants.FrontendChallengeURL, http.StatusTemporaryRedirect)
		})
	})

	app.Route("/:token", func(r fiber.Router) {
		r.Get("/prompt", service.Prompt)
		r.Post("/submit", service.Submit)
	})

	app.Route("/frontend", func(r fiber.Router) {
		r.Get("/movies", service.Frontend)
	})
}

func x404(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotFound)
}
