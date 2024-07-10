package handlers

import (
	"fmt"

	"movie-matcher/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type forgotPromptResponse struct {
	Prompt model.Prompt `json:"prompt"`
}

func (s *Service) ForgotPrompt(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid token %s", rawToken))
	}

	prompt, err := s.storage.ForgotPrompt(c.UserContext(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get prompt %s", err))
	}

	return c.
		Status(fiber.StatusOK).
		JSON(
			forgotPromptResponse{
				Prompt: *prompt,
			},
		)
}
