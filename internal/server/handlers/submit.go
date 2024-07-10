package handlers

import (
	"fmt"

	"movie-matcher/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type submitRequestBody model.Ranking

func (s *Service) Submit(c *fiber.Ctx) error {
	rawToken := c.Params("token")

	token, err := uuid.Parse(rawToken)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("invalid token %s", rawToken))
	}

	var submitRequestBody submitRequestBody
	if err := c.BodyParser(&submitRequestBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("invalid request body %s", submitRequestBody))
	}

	// MARK: @Jackson how to score a submission?
	score := model.Score{}
	if err := s.storage.Submit(c.UserContext(), token, score); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("failed to submit %s", err))
	}

	return c.
		Status(fiber.StatusOK).
		JSON(score)
}
