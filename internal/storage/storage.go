package storage

import (
	"context"
	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/set"

	"github.com/google/uuid"
)

type Storage interface {
	Register(ctx context.Context, nuid applicant.NUID, name applicant.ApplicantName, token uuid.UUID, prompt algo.Prompt, solution set.OrderedSet[movie.ID]) error
	Token(ctx context.Context, nuid applicant.NUID) (*uuid.UUID, error)
	Prompt(ctx context.Context, token uuid.UUID) (*algo.Prompt, error)
	Submit(ctx context.Context, token uuid.UUID, score uint) error
}
