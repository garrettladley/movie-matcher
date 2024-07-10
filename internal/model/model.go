package model

import "time"

type Preference[T any] struct {
	Name   string
	Value  T
	Weight int
}

type Person struct {
	Name           string
	Duration       Preference[time.Duration] `json:"duration"` // refresh myself on omitempty
	Language       Preference[string]
	Actor          Preference[string]
	RottenTomatoes Preference[int]
	Year           Preference[int]
	Rating         Preference[string]
}

type MovieID string // imdb id

var AVAILABLE_MOVIES = []MovieID{
	// jackson select 10-15 imdb id
	// g select 10-15
}

type Ranking struct {
	Movies []MovieID `json:"movies"`
}

type Prompt struct {
	People []Person
	Movies []MovieID
}

type Score struct {
	KendallTau int
}

type MoviePromptService interface {
	Generate(movies []MovieID) Prompt // pass a subset of all available movies
	Check(prompt Prompt, ranking Ranking) Score
}
