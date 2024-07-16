package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Service) Prompt(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	prompt, err := s.storage.Prompt(c.Context(), token)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(prompt)
}
