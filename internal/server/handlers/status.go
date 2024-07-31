package handlers

import (
	"fmt"
	"strings"
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
	email, err := applicant.ParseNUEmail(strings.TrimSpace(rawEmail))
	if err != nil {
		return utilities.IntoTempl(
			c,
			not_found.NotFound(
				not_found.Params{Email: rawEmail},
				not_found.Errors{Email: err.Error()}),
		)
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
			if utilities.IsNotFound(err) {
				// TODO: render a button to take to the docs section on how to register?
				return utilities.IntoTempl(
					c,
					not_found.NotFound(
						not_found.Params{Email: rawEmail},
						not_found.Errors{Email: fmt.Sprintf("user with email %s not found, have you registered?", rawEmail)}),
				)
			}
			return err
		}
	}

	return utilities.IntoTempl(c, status.Index(intoTimePoints(<-submissionsCh), <-nameCh, limit))
}
