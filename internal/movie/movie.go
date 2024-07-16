package movie

import (
	"movie-matcher/internal/duration"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/utilities"
)

type Movie struct {
	ID                  ID
	Year                uint
	AgeRating           string
	Duration            duration.Duration
	Genres              []string
	Directors           []string
	Actors              []string
	Plot                []string
	RottenTomatoesScore uint
}

type ID string // IMDb id

func FromOMDB(omdbMovie omdb.Movie) Movie {
	return Movie{
		ID:                  ID(omdbMovie.IMDbID),
		Year:                omdbMovie.Year,
		AgeRating:           omdbMovie.AgeRating,
		Duration:            omdbMovie.Duration,
		Genres:              omdbMovie.Genres,
		Directors:           omdbMovie.Directors,
		Actors:              omdbMovie.Actors,
		Plot:                utilities.Tokenize(omdbMovie.Plot),
		RottenTomatoesScore: omdbMovie.RottenTomatoesScore,
	}
}
