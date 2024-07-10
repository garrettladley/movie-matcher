package storage

import (
	"context"
	"database/sql"
	"fmt"

	"movie-matcher/internal/model"

	go_json "github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	Register(ctx context.Context, nuid model.NUID, name model.ApplicantName, token uuid.UUID, prompt model.Prompt, solution model.Ranking) error
	ForgotToken(ctx context.Context, nuid model.NUID) (*uuid.UUID, error)
	ForgotPrompt(ctx context.Context, token uuid.UUID) (*model.Prompt, error)
	Submit(ctx context.Context, token uuid.UUID, score model.Score) error
}

type DB struct {
	*sqlx.DB
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{db}
}

func (db *DB) Register(ctx context.Context, nuid model.NUID, name model.ApplicantName, token uuid.UUID, prompt model.Prompt, solution model.Ranking) error {
	marshalledPrompt, err := go_json.Marshal(prompt)
	if err != nil {
		return fmt.Errorf("failed to marshal prompt: %w", err)
	}

	marshalledRanking, err := go_json.Marshal(solution)
	if err != nil {
		return fmt.Errorf("failed to marshal solution: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"INSERT INTO applicants (nuid, applicant_name, token, prompt, solution) VALUES ($1, $2, $3, $4, $5);",
		nuid,
		name,
		token,
		marshalledPrompt,
		marshalledRanking,
	)

	return err
}

func (db *DB) ForgotToken(ctx context.Context, nuid model.NUID) (*uuid.UUID, error) {
	var dbResult struct {
		Token sql.NullString `db:"token"`
	}

	if err := db.GetContext(ctx, &dbResult, "SELECT token FROM applicants WHERE nuid=$1;", nuid); err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	if !dbResult.Token.Valid {
		return nil, fmt.Errorf("token not found")
	}

	token, err := uuid.Parse(dbResult.Token.String)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return &token, nil
}

func (db *DB) ForgotPrompt(ctx context.Context, token uuid.UUID) (*model.Prompt, error) {
	var dbResult struct {
		Prompt sql.NullString `db:"prompt"`
	}

	if err := db.GetContext(
		ctx,
		&dbResult,
		"SELECT prompt FROM applicants WHERE token=$1;",
		token,
	); err != nil {
		return nil, fmt.Errorf("failed to get prompt: %w", err)
	}

	if !dbResult.Prompt.Valid {
		return nil, fmt.Errorf("prompt not found")
	}

	var prompt model.Prompt
	if err := go_json.Unmarshal([]byte(dbResult.Prompt.String), &prompt); err != nil {
		return nil, fmt.Errorf("failed to unmarshal prompt: %w", err)
	}

	return &prompt, nil
}

func (db *DB) Submit(ctx context.Context, token uuid.UUID, score model.Score) error {
	marshalledScore, err := go_json.Marshal(score)
	if err != nil {
		return fmt.Errorf("failed to marshal score: %w", err)
	}

	_, err = db.ExecContext(
		ctx,
		"INSERT INTO submissions (submission_id, token, score) VALUES ($1, $2, $3);",
		uuid.New(),
		token,
		marshalledScore,
	)

	return err
}
