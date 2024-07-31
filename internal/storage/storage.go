package storage

import (
	"context"
	"time"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/model"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/ordered_set"

	"github.com/google/uuid"
)

type Storage interface {
	Register(ctx context.Context, email applicant.NUEmail, name applicant.Name, createdAt time.Time, token uuid.UUID, prompt algo.Prompt, solution ordered_set.OrderedSet[movie.ID]) error
	Token(ctx context.Context, email applicant.NUEmail) (uuid.UUID, error)
	Name(ctx context.Context, token uuid.UUID) (applicant.Name, error)
	Status(ctx context.Context, token uuid.UUID, limit int) ([]model.Submission, error)
	Prompt(ctx context.Context, token uuid.UUID) (algo.Prompt, error)
	Solution(ctx context.Context, token uuid.UUID) (ordered_set.OrderedSet[movie.ID], error)
	Submit(ctx context.Context, token uuid.UUID, score int) error
}
