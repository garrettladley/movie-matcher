package handlers

import (
	"net/http"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/movie"
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
		return utilities.InvalidJSON(err)
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
	prompt := s.moviePrompter.Generate(utilities.SelectRandom(movie.Catalog, 15))
	// MARK: @Jackson how to generate a solution?
	solution := algo.Ranking{Movies: prompt.Movies}

	if err := s.storage.Register(
		c.UserContext(),
		nuid,
		*applicantName,
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
