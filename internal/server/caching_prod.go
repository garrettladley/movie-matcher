//go:build !dev

package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/memory/v2"
)

func setupCaching(app *fiber.App) {
	cacheStorage := memory.New()
	keyGenerator := func(c *fiber.Ctx) string { return utils.CopyString(c.OriginalURL()) }

	app.Use(cache.New(cache.Config{
		Next:         createCacheNextFunction(cacheStorage, keyGenerator, StaticPaths),
		Storage:      cacheStorage,
		KeyGenerator: keyGenerator,
		Expiration:   time.Hour * 24 * 365, // 1 year
		CacheControl: true,
	}))
}

func createCacheNextFunction(storage *memory.Storage, keyGenerator func(c *fiber.Ctx) string, staticPaths map[string]struct{}) func(*fiber.Ctx) bool {
	return func(c *fiber.Ctx) bool {
		if _, ok := staticPaths[c.OriginalURL()]; ok {
			return false
		}

		key := fmt.Sprintf("%s_%s", keyGenerator(c), c.Method())
		value, err := storage.Get(key)
		cacheHit := err == nil && value != nil

		if cacheHit {
			time.Sleep(500 * time.Millisecond)
			return true
		}

		if strings.HasPrefix(c.OriginalURL(), "/frontend/movies") {
			return false
		}

		return true
	}
}
