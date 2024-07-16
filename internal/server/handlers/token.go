package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Token(c *fiber.Ctx) error {
	rawNUID := c.Params("nuid")
	nuid, err := applicant.ParseNUID(rawNUID)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse nuid. got: %s", rawNUID))
	}

	token, err := s.storage.Token(c.Context(), nuid)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		SendString(token.String())
}
