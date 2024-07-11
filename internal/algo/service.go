package algo

import "movie-matcher/internal/model"

type MoviePromptService interface {
	Generate(movies []model.Movie) model.Prompt
	Check(prompt model.Prompt, ranking model.Ranking) model.Score
}

type MoviePrompter struct{}

func (m *MoviePrompter) Generate(movies []model.Movie) model.Prompt {
	return model.Prompt{}
}

func (m *MoviePrompter) Check(prompt model.Prompt, ranking model.Ranking) model.Score {
	return model.Score{}
}
