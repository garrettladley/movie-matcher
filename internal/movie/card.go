package movie

import (
	"movie-matcher/internal/duration"
	"movie-matcher/internal/services/omdb"
)

type Card struct {
	ID                  ID                `json:"id"`
	Year                uint              `json:"year"`
	AgeRating           string            `json:"ageRating"`
	Duration            duration.Duration `json:"duration"`
	Genres              []string          `json:"genres"`
	Directors           []string          `json:"directors"`
	Actors              []string          `json:"actors"`
	Plot                string            `json:"plot"`
	RottenTomatoesScore uint              `json:"rottenTomatoesScore"`
	PosterURL           string            `json:"posterURL"`
	MinutesRemaining    *int              `json:"minutesRemaining,omitempty"`
}

func MovieToCard(omdbMovie omdb.Movie, minutesRemaining *int) Card {
	return Card{
		ID:                  ID(omdbMovie.IMDbID),
		Year:                omdbMovie.Year,
		AgeRating:           omdbMovie.AgeRating,
		Duration:            omdbMovie.Duration,
		Genres:              omdbMovie.Genres,
		Directors:           omdbMovie.Directors,
		Actors:              omdbMovie.Actors,
		Plot:                omdbMovie.Plot,
		RottenTomatoesScore: omdbMovie.RottenTomatoesScore, PosterURL: omdbMovie.Poster,
		MinutesRemaining: minutesRemaining,
	}
}
