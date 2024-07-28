package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/data"
	"movie-matcher/internal/server/ctxt"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Chart(c *fiber.Ctx) error {
	rawNUID := c.Params("nuid")
	nuid, err := applicant.ParseNUID(rawNUID)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse nuid. got: %s", rawNUID))
	}

	ctxt.WithNUID(c, nuid)

	limit := c.QueryInt("limit", 5)

	submissions, err := s.storage.Status(c.Context(), nuid, limit)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(data.Into(intoTimePoints(submissions)))
}
