package handlers

import (
	"movie-matcher/internal/algo"
	"movie-matcher/internal/storage"
)

type Service struct {
	storage storage.Storage
	algo    *algo.Service
}

func NewService(storage storage.Storage, algo *algo.Service) *Service {
	return &Service{
		storage: storage,
		algo:    algo,
	}
}
