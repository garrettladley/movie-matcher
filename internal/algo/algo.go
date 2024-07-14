package algo

import (
	"movie-matcher/internal/movie"
	"movie-matcher/internal/services/pref_gen"
)

type Ranking struct {
	Movies []movie.ID `json:"movies"`
}

type Prompt struct {
	People []pref_gen.Person `json:"people"`
	Movies []movie.ID        `json:"movies"`
}

type Score struct {
	KendallTau int `json:"kendall_tau"`
}
