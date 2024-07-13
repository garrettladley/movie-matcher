package movie

import "time"

type Movie struct {
	Title               string
	ID                  ID
	Year                uint
	AgeRating           Rating
	Duration            time.Duration
	Languages           []Language
	Genres              []string
	Directors           []string
	Writers             []string
	Actors              []string
	Plot                string
	IMDbScore           uint
	RottenTomatoesScore uint
	MetacriticScore     uint
}

type ID string // IMDb id

type MovieBuilder struct {
	movie Movie
}

func NewMovieBuilder() *MovieBuilder {
	return &MovieBuilder{}
}

func (m *MovieBuilder) ID(id ID) *MovieBuilder {
	m.movie.ID = id
	return m
}

func (m *MovieBuilder) Runtime(duration time.Duration) *MovieBuilder {
	m.movie.Duration = duration
	return m
}

func (m *MovieBuilder) Languages(languages []Language) *MovieBuilder {
	m.movie.Languages = languages
	return m
}

func (m *MovieBuilder) Actors(actors []string) *MovieBuilder {
	m.movie.Actors = actors
	return m
}

func (m *MovieBuilder) RottenTomatoes(score uint) *MovieBuilder {
	m.movie.RottenTomatoesScore = score
	return m
}

func (m *MovieBuilder) Year(year uint) *MovieBuilder {
	m.movie.Year = year
	return m
}

func (m *MovieBuilder) Rating(rating Rating) *MovieBuilder {
	m.movie.AgeRating = rating
	return m
}

func (m *MovieBuilder) Build() Movie {
	return m.movie
}
