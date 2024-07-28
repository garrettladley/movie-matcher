package handlers

import (
	"fmt"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/server/ctxt"
	"movie-matcher/internal/utilities"
	"movie-matcher/internal/views/status"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Status(c *fiber.Ctx) error {
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

	name, err := s.storage.Name(c.Context(), nuid)
	if err != nil {
		return err
	}

	return into(c, status.Index(intoTimePoints(submissions), name, limit))
}
