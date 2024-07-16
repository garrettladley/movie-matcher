package handlers

import (
	"fmt"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/utilities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Frontend(c *fiber.Ctx) error {
	fetch_type := c.Query("type")
	if fetch_type == "" {
		return utilities.BadRequest(fmt.Errorf("missing type query parameter"))
	}

	var movies []movie.ID
	watching := false
	switch fetch_type {
		case "top": {
			movies = movie.TopMoviesCatalog
		}
		case "recommended": {
			movies = movie.RecommendationsCatalog
		}
		case "watching": {
			movies = movie.ContinueWatchingCatalog
			watching = true
		}
		default: {
			return utilities.BadRequest(fmt.Errorf("invalid type query parameter"))
		}
	}

	// Loop over each movie and fetch content from OMDB
	var moviesData []movie.MovieDisplay
	for i := 0 ; i < len(movies) ; i++ {
		movieData, err := s.algo.Client.FindMovieById(c.UserContext(), string(movies[i]))
		if err != nil {
			return err
		}

		movieDisplay, err := movie.MovieToDisplay(movie.FromOMDB(movieData), watching)
		if err != nil {
			return err
		}

		moviesData = append(moviesData, *movieDisplay)
	}

	return c.Status(http.StatusOK).JSON(moviesData)
}