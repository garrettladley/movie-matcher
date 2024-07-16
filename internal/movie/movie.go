package movie

import (
	"movie-matcher/internal/duration"
	"movie-matcher/internal/services/omdb"
)

type Movie struct {
	Title               string
	ID                  ID
	Year                uint
	AgeRating           string
	Duration            duration.Duration
	Languages           []string
	Genres              []string
	Directors           []string
	Writers             []string
	Actors              []string
	Plot                []string
	IMDbScore           uint
	RottenTomatoesScore uint
	MetacriticScore     uint
}

type ID string // IMDb id

func FromOMDB(omdbMovie omdb.Movie) Movie {
	return Movie{
		Title:               omdbMovie.Title,
		ID:                  ID(omdbMovie.IMDbID),
		Year:                omdbMovie.Year,
		AgeRating:           omdbMovie.AgeRating,
		Duration:            omdbMovie.Duration,
		Languages:           omdbMovie.Languages,
		Genres:              omdbMovie.Genres,
		Directors:           omdbMovie.Directors,
		Writers:             omdbMovie.Writers,
		Actors:              omdbMovie.Actors,
		Plot:                omdbMovie.Plot,
		IMDbScore:           omdbMovie.IMDbScore,
		RottenTomatoesScore: omdbMovie.RottenTomatoesScore,
		MetacriticScore:     omdbMovie.MetacriticScore,
	}
}
