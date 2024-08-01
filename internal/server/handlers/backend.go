package handlers

import (
	"movie-matcher/internal/views/backend"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Backend(c *fiber.Ctx) error {
	return into(c, backend.Backend())
}
