package handlers

import (
	"movie-matcher/internal/algo"
	"movie-matcher/internal/storage"
)

type Service struct {
	storage       storage.Storage
	moviePrompter algo.MoviePromptService
}

func NewService(storage storage.Storage, moviePrompter algo.MoviePromptService) *Service {
	return &Service{
		storage:       storage,
		moviePrompter: moviePrompter,
	}
}
