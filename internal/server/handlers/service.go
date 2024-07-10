package handlers

import (
	"movie-matcher/internal/model"
	"movie-matcher/internal/storage"
)

type Service struct {
	storage       storage.Storage
	moviePrompter model.MoviePromptService
}

func NewService(storage storage.Storage, moviePrompter model.MoviePromptService) *Service {
	return &Service{
		storage:       storage,
		moviePrompter: moviePrompter,
	}
}
