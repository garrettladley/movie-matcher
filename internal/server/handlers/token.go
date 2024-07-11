package handlers

import (
	"fmt"
	"net/http"

	"movie-matcher/internal/model"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TokenResponse struct {
	Token uuid.UUID `json:"token"`
}

func (s *Service) Token(c *fiber.Ctx) error {
	rawNUID := c.Params("nuid")
	nuid, err := model.ParseNUID(rawNUID)
	if err != nil {
		return utilities.BadRequest(fmt.Errorf("failed to parse nuid. got: %s", rawNUID))
	}

	token, err := s.storage.Token(c.UserContext(), *nuid)
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(
			TokenResponse{
				Token: *token,
			},
		)
}
