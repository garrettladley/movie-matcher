package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"movie-matcher/internal/config"
	"movie-matcher/internal/server"
)

func main() {
	settings, err := config.GetSettings("config")
	if err != nil {
		panic(err)
	}

	app := server.Setup(*settings)

	go func() {
		if err := app.Listen(strconv.Itoa(int(settings.Application.Port))); err != nil {
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
