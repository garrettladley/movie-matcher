package model

import "time"

type MovieID string // imdb id

type Movie struct {
	ID             MovieID       `json:"id"`
	Runtime        time.Duration `json:"runtime"`
	Languages      []Language    `json:"languages"`
	Actors         []string      `json:"actors"`
	RottenTomatoes int           `json:"rotten_tomatoes"`
	Year           int           `json:"year"`
	Rating         Rating        `json:"rating"`
}

type MovieBuilder struct {
	movie Movie
}

func NewMovieBuilder() *MovieBuilder {
	return &MovieBuilder{}
}

func (m *MovieBuilder) ID(id MovieID) *MovieBuilder {
	m.movie.ID = id
	return m
}

func (m *MovieBuilder) Runtime(runtime time.Duration) *MovieBuilder {
	m.movie.Runtime = runtime
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

func (m *MovieBuilder) RottenTomatoes(rottenTomatoes int) *MovieBuilder {
	m.movie.RottenTomatoes = rottenTomatoes
	return m
}

func (m *MovieBuilder) Year(year int) *MovieBuilder {
	m.movie.Year = year
	return m
}

func (m *MovieBuilder) Rating(rating Rating) *MovieBuilder {
	m.movie.Rating = rating
	return m
}

func (m *MovieBuilder) Build() Movie {
	return m.movie
}
