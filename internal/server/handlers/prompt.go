package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PromptResponse struct {
	Prompt algo.Prompt `json:"prompt"`
}

func (s *Service) Prompt(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	prompt, err := s.storage.Prompt(c.UserContext(), token)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(
			PromptResponse{
				Prompt: *prompt,
			},
		)
}
