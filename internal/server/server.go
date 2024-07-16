package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"movie-matcher/internal/config"
	"movie-matcher/internal/server/handlers"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/storage"
	"movie-matcher/internal/utilities"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/memory/v2"

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

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	service := handlers.NewService(
		storage.NewPostgresDB(settings.Database),
		omdb.NewCachedClient(),
	)

	cacheStorage := memory.New()
	keyGenerator := func(c *fiber.Ctx) string { return utils.CopyString(c.OriginalURL()) }

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			key := fmt.Sprintf("%s_%s", keyGenerator(c), c.Method())
			value, err := cacheStorage.Get(key)
			cacheHit := err == nil && value != nil

			if cacheHit {
				time.Sleep(500 * time.Millisecond)
				return true
			}

			if strings.HasPrefix(c.OriginalURL(), "/frontend/movies") {
				return false
			}

			return true
		},
		Storage:      cacheStorage,
		KeyGenerator: keyGenerator,
		Expiration:   time.Hour * 24 * 365, // 1 year
	}))

	app.Route("/",
		func(r fiber.Router) {
			r.Post("register", service.Register)
			r.Route(":nuid", func(r fiber.Router) {
				r.Get("token", service.Token)
			})
			r.Route(":token", func(r fiber.Router) {
				r.Get("prompt", service.Prompt)
				r.Post("submit", service.Submit)
			})
			r.Route("frontend", func(r fiber.Router) {
				r.Get("movies", service.Frontend)
			})
		},
	)

	return app
}
