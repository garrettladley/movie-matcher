package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/data"
	"movie-matcher/internal/server/ctxt"

	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Service) Chart(c *fiber.Ctx) error {
	rawToken := c.Query("token")

	token, err := uuid.Parse(rawToken)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse token. got: %s", token))
	}

	ctxt.WithToken(c, token)

	limit := c.QueryInt("limit", 5)

	submissions, err := s.storage.Status(c.Context(), token, limit)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(data.Into(intoTimePoints(submissions)))
}
