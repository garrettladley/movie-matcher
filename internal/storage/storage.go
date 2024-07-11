package storage

import (
	"context"

	"movie-matcher/internal/model"

	"github.com/google/uuid"
)

type Storage interface {
	Register(ctx context.Context, nuid model.NUID, name model.ApplicantName, token uuid.UUID, prompt model.Prompt, solution model.Ranking) error
	Token(ctx context.Context, nuid model.NUID) (*uuid.UUID, error)
	Prompt(ctx context.Context, token uuid.UUID) (*model.Prompt, error)
	Submit(ctx context.Context, token uuid.UUID, score model.Score) error
}
