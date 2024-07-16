package movie

import (
	"movie-matcher/internal/services/omdb"
)

type Card struct {
	common
	Plot             string `json:"plot"`
	PosterURL        string `json:"posterURL"`
	MinutesRemaining *int   `json:"minutesRemaining,omitempty"`
}

func MovieToCard(omdbMovie omdb.Movie, minutesRemaining *int) Card {
	return Card{
		common:           *commonFrom(omdbMovie),
		Plot:             omdbMovie.Plot,
		PosterURL:        omdbMovie.Poster,
		MinutesRemaining: minutesRemaining,
	}
}
