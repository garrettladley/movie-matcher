package movie

import (
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/utilities"
)

type Movie struct {
	common
	Plot []string
}

type ID string // IMDb id

func FromOMDB(omdbMovie omdb.Movie) Movie {
	return Movie{
		common: *commonFrom(omdbMovie),
		Plot:   utilities.Tokenize(omdbMovie.Plot),
	}
}
