package handlers

import (
	"sync"

	"movie-matcher/internal/applicant"
	"movie-matcher/internal/model"
	"movie-matcher/internal/server/ctxt"
	"movie-matcher/internal/utilities"
	"movie-matcher/internal/views/status"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Service) Status(c *fiber.Ctx) error {
	rawToken := c.Query("token")
	token, err := uuid.Parse(rawToken)

	var params status.StatusParams[int]
	params.Token = rawToken

	var errs status.StatusErrors

	if err != nil {
		errs.Token = "The token provided does not match the expected format."
		return utilities.IntoTempl(c, status.Index(params, errs))
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
			errs.Token = "The provided token is invalid."
			return utilities.IntoTempl(c, status.Index(params, errs))
		}
	}
	
	params.Token = token.String()
	params.Timeseries = intoTimePoints(<-submissionsCh)
	params.Name = <-nameCh
	params.CurrentLimit = limit

	return utilities.IntoTempl(c, status.Index(params, errs))
}
