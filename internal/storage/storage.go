package storage

import (
	"context"
	"time"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/ordered_set"

	"github.com/google/uuid"
)

type Storage interface {
	Register(ctx context.Context, nuid applicant.NUID, name applicant.ApplicantName, createdAt time.Time, token uuid.UUID, prompt algo.Prompt, solution ordered_set.OrderedSet[movie.ID]) error
	Token(ctx context.Context, nuid applicant.NUID) (uuid.UUID, error)
	Prompt(ctx context.Context, token uuid.UUID) (algo.Prompt, error)
	Solution(ctx context.Context, token uuid.UUID) (ordered_set.OrderedSet[movie.ID], error)
	Submit(ctx context.Context, token uuid.UUID, score int) error
}
