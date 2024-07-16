package movie

import (
	"movie-matcher/internal/duration"
	"movie-matcher/internal/services/omdb"
)

type common struct {
	ID                  ID
	Year                uint
	AgeRating           string
	Duration            duration.Duration
	Genres              []string
	Directors           []string
	Actors              []string
	RottenTomatoesScore uint
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
