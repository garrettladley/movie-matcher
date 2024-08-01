//go:build dev

package server

import "github.com/gofiber/fiber/v2"

func setupCaching(app *fiber.App) map[string]struct{} {
	return map[string]struct{}{
		"/deps/apexcharts.min.js": {},
		"/deps/htmx.min.js":       {},
		"/deps/flowbite.min.js":   {},
		"/public/styles.css":      {},
	}
}
