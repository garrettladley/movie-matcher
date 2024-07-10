package handlers

import (
	"fmt"

	"movie-matcher/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type forgotTokenResponse struct {
	Token uuid.UUID `json:"token"`
}

func (s *Service) ForgotToken(c *fiber.Ctx) error {
	rawNUID := c.Params("nuid")
	nuid, err := model.ParseNUID(rawNUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid NUID %s", rawNUID))
	}

	token, err := s.storage.ForgotToken(c.UserContext(), *nuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get token %s", err))
	}

	return c.
		Status(fiber.StatusOK).
		JSON(
			forgotTokenResponse{
				Token: *token,
			},
		)
}
