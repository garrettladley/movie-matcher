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
)

func (s *Service) Status(c *fiber.Ctx) error {
	rawEmail := c.Query("email")
	email, err := applicant.ParseNUEmail(rawEmail)

	notFoundParams := not_found.NotFoundParams{
		Email: rawEmail,
	}
	var notFoundErrs not_found.NotFoundErrors

	if err != nil {
		notFoundErrs.Email = "The email provided is not a valid @northeastern.edu address"
		return utilities.IntoTempl(c, not_found.Index(notFoundParams, notFoundErrs))
	}

	ctxt.WithEmail(c, email)

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
		submissions, err := s.storage.Status(c.Context(), email, limit)
		if err != nil {
			errCh <- err
			return
		}
		submissionsCh <- submissions
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		name, err := s.storage.Name(c.Context(), email)
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
			notFoundErrs.Email = "The provided email is not registered."
			return utilities.IntoTempl(c, not_found.Index(notFoundParams, notFoundErrs))
		}
	}
	
	params := status.StatusParams[int]{
		Email: email,
		Timeseries: intoTimePoints(<-submissionsCh),
		Name: <-nameCh,
		CurrentLimit: limit,
	}

	return utilities.IntoTempl(c, status.Index(params))
}
