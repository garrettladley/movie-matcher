package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"movie-matcher/internal/movie"
	"movie-matcher/internal/set"
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

	prompt, err := s.storage.Prompt(c.UserContext(), token)
	if err != nil {
		return err
	}

	score, err := s.algo.Check(c.UserContext(), *prompt, set.NewOrderedSet(submitRequestBody...))
	if err != nil {
		return err
	}

	if err := s.storage.Submit(c.UserContext(), token, score); err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(score)
}
