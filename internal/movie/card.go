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
	common := *commonFrom(omdbMovie)
	return Card{
		common:           common,
		Plot:             omdbMovie.Plot,
		PosterURL:        PosterCatalog[common.ID],
		MinutesRemaining: minutesRemaining,
	}
}
