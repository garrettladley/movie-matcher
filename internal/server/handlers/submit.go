package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/movie"
	"movie-matcher/internal/set"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type submitRequestBody set.OrderedSet[movie.ID]

func (s *Service) Submit(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	var submitRequestBody submitRequestBody
	if err := c.BodyParser(&submitRequestBody); err != nil {
		return utilities.InvalidJSON(err)
	}

	// MARK: @Jackson how to score a submission?
	var score uint = 0
	if err := s.storage.Submit(c.UserContext(), token, score); err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(score)
}
