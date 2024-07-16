package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"movie-matcher/internal/movie"
	"movie-matcher/internal/ordered_set"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type submitRequestBody []movie.ID

func (s *Service) Submit(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	var submitRequestBody submitRequestBody
	if err := c.BodyParser(&submitRequestBody); err != nil {
		slog.Error("invalid JSON request data", "error", err)
		return utilities.InvalidJSON()
	}

	prompt, err := s.storage.Prompt(c.Context(), token)
	if err != nil {
		return err
	}

	score, err := s.algo.Check(c.Context(), *prompt, ordered_set.New(submitRequestBody...))
	if err != nil {
		return err
	}

	if err := s.storage.Submit(c.Context(), token, score); err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(score)
}
