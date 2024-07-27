package main

import (
	"embed"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"movie-matcher/internal/config"
	"movie-matcher/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	settings, err := config.GetSettings("config")
	if err != nil {
		log.Fatal(err)
	}

	app := server.Setup(*settings)

	static(app)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", settings.Application.Port)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("Server shutdown")
}

//go:embed public
var PublicFS embed.FS

//go:embed htmx
var HtmxFS embed.FS

func static(app *fiber.App) {
	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(PublicFS),
		PathPrefix: "public",
		Browse:     true,
	}))
	app.Use("/htmx", filesystem.New(filesystem.Config{
		Root:       http.FS(HtmxFS),
		PathPrefix: "htmx",
		Browse:     true,
	}))
}
