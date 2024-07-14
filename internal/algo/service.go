package algo

import (
	"movie-matcher/internal/movie"
)

type MoviePromptService interface {
	Generate(movies []movie.Movie) Prompt
	Check(prompt Prompt, ranking Ranking) Score
}

type MoviePrompter struct{}

func (m *MoviePrompter) Generate(movies []movie.Movie) Prompt {
	return Prompt{}
}

func (m *MoviePrompter) Check(prompt Prompt, ranking Ranking) Score {
	return Score{}
}
