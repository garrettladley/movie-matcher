package handlers

import (
	"sync"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/model"
	"movie-matcher/internal/server/ctxt"
	"movie-matcher/internal/utilities"
	"movie-matcher/internal/views/not_found"
	"movie-matcher/internal/views/status"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Service) Status(c *fiber.Ctx) error {
	rawToken := c.Query("token")
	token, err := uuid.Parse(rawToken)

	notFoundParams := not_found.NotFoundParams{
		Token: rawToken,
	}
	var notFoundErrs not_found.NotFoundErrors

	var errs status.StatusErrors

	if err != nil {
		notFoundErrs.Token = "The token provided does not match the expected format."
		return utilities.IntoTempl(c, not_found.Index(notFoundParams, notFoundErrs))
	}

	ctxt.WithToken(c, token)

	limit := c.QueryInt("limit", 5)

	var (
		submissionsCh = make(chan []model.Submission, 1)
		nameCh        = make(chan applicant.Name, 1)
		errCh         = make(chan error, 2)

		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		submissions, err := s.storage.Status(c.Context(), token, limit)
		if err != nil {
			errCh <- err
			return
		}
		submissionsCh <- submissions
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		name, err := s.storage.Name(c.Context(), token)
		if err != nil {
			errCh <- err
			return
		}
		nameCh <- name
	}()

	wg.Wait()
	close(submissionsCh)
	close(nameCh)
	close(errCh)

	for err := range errCh {
		if err != nil {
			notFoundErrs.Token = "The provided token is invalid."
			return utilities.IntoTempl(c, not_found.Index(notFoundParams, notFoundErrs))
		}
	}
	
	params := status.StatusParams[int]{
		Token: token.String(),
		Timeseries: intoTimePoints(<-submissionsCh),
		Name: <-nameCh,
		CurrentLimit: limit,
	}

	return utilities.IntoTempl(c, status.Index(params, errs))
}
