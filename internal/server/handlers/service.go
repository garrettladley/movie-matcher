package handlers

import (
	"movie-matcher/internal/algo"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/storage"
)

type Service struct {
	storage storage.Storage
	client  *omdb.CachedClient
	algo    *algo.Service
}

func NewService(storage storage.Storage, client *omdb.CachedClient) *Service {
	return &Service{
		storage: storage,
		client:  client,
		algo:    algo.NewService(client),
	}
}
