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
	"movie-matcher/internal/views/not_found"

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
		JSONEncoder:       go_json.Marshal,
		JSONDecoder:       go_json.Unmarshal,
		ErrorHandler:      utilities.ErrorHandler,
		PassLocalsToViews: true,
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

	staticPaths := map[string]struct{}{
		"/deps/apexcharts.min.js": {},
		"/deps/htmx.min.js":       {},
		"/deps/flowbite.min.js":   {},
		"/public/styles.css":      {},
	}

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			if _, ok := staticPaths[c.OriginalURL()]; ok {
				return false
			}

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
		CacheControl: true,
	}))

	app.Route("/",
		func(r fiber.Router) {
			r.Get("favicon.ico", x404)
			r.Get("/deps/flowbite.min.js.map", x404)
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
		},
	)

	app.Use(func(c *fiber.Ctx) error {
		if _, ok := staticPaths[c.OriginalURL()]; ok {
			return c.Next()
		}
		return utilities.IntoTempl(c, not_found.NotFound(not_found.Params{}, not_found.Errors{}))
	})

	return app
}

func x404(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotFound)
}
