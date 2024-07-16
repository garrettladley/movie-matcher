package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"movie-matcher/internal/movie"
	"movie-matcher/internal/ordered_set"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Frontend(c *fiber.Ctx) error {
	var (
		fetchType string = c.Query("type")
		movies    ordered_set.OrderedSet[movie.ID]
		watching  bool
	)

	switch fetchType {
	case "top":
		movies = movie.TopMoviesCatalog
	case "recommended":
		movies = movie.RecommendationsCatalog
	case "watching":
		movies = movie.ContinueWatchingCatalog
		watching = true
	default:
		return utilities.BadRequest(fmt.Errorf("invalid type query parameter. got: '%s'", fetchType))
	}

	cards, err := s.fetchFrontendMovies(c.Context(), movies, watching)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(cards)
}

func (s *Service) fetchFrontendMovies(ctx context.Context, movies ordered_set.OrderedSet[movie.ID], watching bool) ([]movie.Card, error) {
	var (
		wg      sync.WaitGroup
		n       uint         = movies.Len()
		cards   []movie.Card = make([]movie.Card, n)
		errChan chan error   = make(chan error, n)
	)

	for index, id := range movies.Slice() {
		wg.Add(1)
		go func(index int, id movie.ID) {
			defer wg.Done()

			m, err := s.client.FindMovieById(ctx, string(id))
			if err != nil {
				errChan <- fmt.Errorf("failed to calculate score for movie %s: %w", id, err)
				return
			}

			var minutesRemaining *int
			if watching {
				randomMinutesRemaining := generateMinutesRemaining(time.Now().UnixNano(), m)
				minutesRemaining = &randomMinutesRemaining
			}

			cards[index] = movie.MovieToCard(m, minutesRemaining)
		}(index, id)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		cards = []movie.Card{}
	}

	return cards, <-errChan
}

func generateMinutesRemaining(seed int64, omdbMovie omdb.Movie) int {
	return rand.New(rand.NewSource(seed)).Intn(int(omdbMovie.Duration.Value().Minutes())) + 1
}
