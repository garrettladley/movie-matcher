package handlers

import (
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type registerRequest struct {
	RawApplicantName string `json:"name"`
	RawNUID          string `json:"nuid"`
}

type registerResponse struct {
	Token  uuid.UUID   `json:"token"`
	Prompt algo.Prompt `json:"prompt"`
}

func (s *Service) Register(c *fiber.Ctx) error {
	var registerRequestBody registerRequest
	if err := c.BodyParser(&registerRequestBody); err != nil {
		slog.Error("invalid JSON request data", "error", err)
		return utilities.InvalidJSON()
	}

	errors := make(map[string]string)
	nuid, err := applicant.ParseNUID(registerRequestBody.RawNUID)
	if err != nil {
		errors["nuid"] = err.Error()
	}

	applicantName, err := applicant.ParseApplicantName(registerRequestBody.RawApplicantName)
	if err != nil {
		errors["name"] = err.Error()
	}

	if len(errors) > 0 {
		return utilities.InvalidRequestData(errors)
	}

	token := uuid.New()
	now := time.Now()
	prompt := s.algo.Generate(rand.New(rand.NewSource(now.UnixNano())))
	solution, err := s.algo.Solution(c.Context(), prompt.Movies, prompt.People)
	if err != nil {
		return err
	}

	if err := s.storage.Register(
		c.Context(),
		nuid,
		applicantName,
		now,
		token,
		prompt,
		solution,
	); err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(
			registerResponse{
				Token:  token,
				Prompt: prompt,
			})
}
