package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/data"
	"movie-matcher/internal/server/ctxt"

	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Chart(c *fiber.Ctx) error {
	rawEmail := c.Query("email")
	email, err := applicant.ParseNUEmail(rawEmail)

	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse email. got: %s", email))
	}

	ctxt.WithEmail(c, email)

	limit := c.QueryInt("limit", 5)

	submissions, err := s.storage.Status(c.Context(), email, limit)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(data.Into(intoTimePoints(submissions)))
}
