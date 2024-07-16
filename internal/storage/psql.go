package storage

import (
	"context"
	"database/sql"
	"time"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/config"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/ordered_set"
	"movie-matcher/internal/utilities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	go_json "github.com/goccy/go-json"
)

type PostgresDB struct {
	*sqlx.DB
}

func NewPostgresDB(settings config.DatabaseSettings) *PostgresDB {
	return &PostgresDB{sqlx.MustConnect("postgres", settings.WithDb())}
}

func (db *PostgresDB) Register(ctx context.Context, nuid applicant.NUID, name applicant.ApplicantName, createdAt time.Time, token uuid.UUID, prompt algo.Prompt, solution ordered_set.OrderedSet[movie.ID]) error {
	marshalledPrompt, err := go_json.Marshal(prompt)
	if err != nil {
		return err
	}

	marshalledRanking, err := go_json.Marshal(solution)
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(
		ctx,
		"INSERT INTO applicants (nuid, applicant_name, created_at, token, prompt, solution) VALUES ($1, $2, $3, $4, $5, $6);",
		nuid,
		name,
		createdAt,
		token,
		marshalledPrompt,
		marshalledRanking,
	); err != nil {
		if db.isUniqueViolation(err) {
			return utilities.Conflict("user", "nuid")
		} else {
			return err
		}
	}

	return nil
}

func (db *PostgresDB) Token(ctx context.Context, nuid applicant.NUID) (*uuid.UUID, error) {
	var dbResult struct {
		Token sql.NullString `db:"token"`
	}

	if err := db.GetContext(ctx, &dbResult, "SELECT token FROM applicants WHERE nuid=$1;", nuid); err != nil {
		return nil, err
	}

	if !dbResult.Token.Valid {
		return nil, utilities.NotFound("token")
	}

	token, err := uuid.Parse(dbResult.Token.String)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (db *PostgresDB) Prompt(ctx context.Context, token uuid.UUID) (*algo.Prompt, error) {
	var dbResult struct {
		Prompt sql.NullString `db:"prompt"`
	}

	if err := db.GetContext(
		ctx,
		&dbResult,
		"SELECT prompt FROM applicants WHERE token=$1;",
		token,
	); err != nil {
		return nil, err
	}

	if !dbResult.Prompt.Valid {
		return nil, utilities.NotFound("prompt")
	}

	var prompt algo.Prompt
	if err := go_json.Unmarshal([]byte(dbResult.Prompt.String), &prompt); err != nil {
		return nil, err
	}

	return &prompt, nil
}

func (db *PostgresDB) Submit(ctx context.Context, token uuid.UUID, score int) error {
	marshalledScore, err := go_json.Marshal(score)
	if err != nil {
		return err
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

func (db *PostgresDB) isUniqueViolation(err error) bool {
	pgErr, isPGError := err.(*pq.Error)
	return isPGError && pgErr.Code == "23505"
}
