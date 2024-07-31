package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Token(c *fiber.Ctx) error {
	rawEmail := c.Query("email")
	email, err := applicant.ParseNUEmail(rawEmail)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse email. got: %s", email))
	}

	token, err := s.storage.Token(c.Context(), email)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		SendString(token.String())
}
