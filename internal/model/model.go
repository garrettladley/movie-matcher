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
	"tt1216475",  // whiplash
	"tt15239678", // dune 2
	"tt17978434", // spider-man: into the spiderverse
	"tt2975976",  // lalaland
	"tt1490017",  // the lego movie
	"tt0062622",  // 2001: a space odyssey
	"tt22022452", // inside out 2
	"tt2278388",  // the grand budapest hotel
	"tt2084970",  // the imitation game
	"tt0112384",  // apollo 13
	"tt0264464",  // catch me if you can
	"tt0432283",  // fantastic mr fox
	"tt2293640",  // the minions movie
	"tt0058150",  // goldfinger
	"tt1074638",  // skyfall
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

type MoviePrompter struct{}

func (m *MoviePrompter) Generate(movies []MovieID) Prompt {
	return Prompt{}
}

func (m *MoviePrompter) Check(prompt Prompt, ranking Ranking) Score {
	return Score{}
}
