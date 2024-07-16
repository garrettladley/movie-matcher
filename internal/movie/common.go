package movie

import (
	"movie-matcher/internal/duration"
	"movie-matcher/internal/services/omdb"
)

type common struct {
	ID                  ID                `json:"id"`
	Year                uint              `json:"year"`
	AgeRating           string            `json:"ageRating"`
	Duration            duration.Duration `json:"duration"`
	Genres              []string          `json:"genres"`
	Directors           []string          `json:"directors"`
	Actors              []string          `json:"actors"`
	RottenTomatoesScore uint              `json:"rottenTomatoesScore"`
}

func commonFrom(omdbMovie omdb.Movie) *common {
	return &common{
		ID:                  ID(omdbMovie.IMDbID),
		Year:                omdbMovie.Year,
		AgeRating:           omdbMovie.AgeRating,
		Duration:            omdbMovie.Duration,
		Genres:              omdbMovie.Genres,
		Directors:           omdbMovie.Directors,
		Actors:              omdbMovie.Actors,
		RottenTomatoesScore: omdbMovie.RottenTomatoesScore,
	}
}
