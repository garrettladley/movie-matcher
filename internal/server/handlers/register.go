package handlers

import (
	"fmt"

	"movie-matcher/internal/model"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type registerRequest struct {
	RawApplicantName string `json:"name"`
	RawNUID          string `json:"nuid"`
}

type registerResponse struct {
	Token  uuid.UUID    `json:"token"`
	Prompt model.Prompt `json:"prompt"`
}

func (s *Service) Register(c *fiber.Ctx) error {
	var registerRequestBody registerRequest
	if err := c.BodyParser(&registerRequestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid request body %s", registerRequestBody))
	}

	nuid, err := model.ParseNUID(registerRequestBody.RawNUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid NUID %s", registerRequestBody.RawNUID))
	}

	applicantName, err := model.ParseApplicantName(registerRequestBody.RawApplicantName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid applicant name %s", registerRequestBody.RawApplicantName))
	}

	token := uuid.New()
	prompt := s.moviePrompter.Generate(utilities.SelectRandom(model.AVAILABLE_MOVIES, 15))
	// MARK: @Jackson how to generate a solution?
	solution := model.Ranking{Movies: prompt.Movies}

	if err := s.storage.Register(
		c.UserContext(),
		*nuid,
		*applicantName,
		token,
		prompt,
		solution,
	); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to register %s", err))
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(
			registerResponse{
				Token:  token,
				Prompt: prompt,
			})
}
