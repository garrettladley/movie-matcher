package handlers

import (
	"movie-matcher/internal/utilities"
	"movie-matcher/internal/views/home"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Home(c *fiber.Ctx) error {
	return utilities.Render(c, home.Index())
}
