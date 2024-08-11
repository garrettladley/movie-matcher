package handlers

import (
	"movie-matcher/internal/utilities"
	"movie-matcher/internal/views/index"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Index(c *fiber.Ctx) error {
	return utilities.Render(c, index.Index())
}
