package storage

import (
	"context"
	"database/sql"
	"time"

	"movie-matcher/internal/algo"
	"movie-matcher/internal/applicant"
	"movie-matcher/internal/config"
	"movie-matcher/internal/model"
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

func (db *PostgresDB) Register(ctx context.Context, email applicant.NUEmail, name applicant.Name, createdAt time.Time, token uuid.UUID, prompt algo.Prompt, solution ordered_set.OrderedSet[movie.ID]) error {
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
		"INSERT INTO applicants (email, applicant_name, created_at, token, prompt, solution) VALUES ($1, $2, $3, $4, $5, $6);",
		email,
		name,
		createdAt,
		token,
		marshalledPrompt,
		marshalledRanking,
	); err != nil {
		if db.isUniqueViolation(err) {
			return utilities.Conflict("user", "email")
		} else {
			return err
		}
	}

	return nil
}

type tokenResult struct {
	Token sql.NullString `db:"token"`
}

func (db *PostgresDB) Token(ctx context.Context, email applicant.NUEmail) (uuid.UUID, error) {
	var dbResult tokenResult
	if err := db.GetContext(ctx, &dbResult, "SELECT token FROM applicants WHERE email=$1;", email); err != nil {
		return uuid.UUID{}, err
	}

	if !dbResult.Token.Valid {
		return uuid.UUID{}, utilities.NotFound("token")
	}

	token, err := uuid.Parse(dbResult.Token.String)
	if err != nil {
		return uuid.UUID{}, err
	}

	return token, nil
}

func (db *PostgresDB) Name(ctx context.Context, token uuid.UUID) (applicant.Name, error) {
	var name applicant.Name
	if err := db.GetContext(ctx, &name, "SELECT applicant_name FROM applicants WHERE token=$1;", token); err != nil {
		return "", err
	}

	return name, nil
}

func (db *PostgresDB) Status(ctx context.Context, token uuid.UUID, limit int) ([]model.Submission, error) {
	var submissions []model.Submission
	query := `
        SELECT s.score, s.submission_time
        FROM submissions s
        INNER JOIN applicants a ON s.token = a.token
        WHERE a.token = $1
        ORDER BY s.submission_time DESC
        LIMIT $2
    `
	if err := db.SelectContext(ctx, &submissions, query, token, limit); err != nil {
		return nil, err
	}

	return submissions, nil
}

type promptResult struct {
	Prompt sql.NullString `db:"prompt"`
}

func (db *PostgresDB) Prompt(ctx context.Context, token uuid.UUID) (algo.Prompt, error) {
	var dbResult promptResult
	if err := db.GetContext(ctx, &dbResult, "SELECT prompt FROM applicants WHERE token=$1;", token); err != nil {
		return algo.Prompt{}, err
	}

	if !dbResult.Prompt.Valid {
		return algo.Prompt{}, utilities.NotFound("prompt")
	}

	var prompt algo.Prompt
	if err := go_json.Unmarshal([]byte(dbResult.Prompt.String), &prompt); err != nil {
		return algo.Prompt{}, err
	}

	return prompt, nil
}

type solutionResult struct {
	Solution sql.NullString `db:"solution"`
}

func (db *PostgresDB) Solution(ctx context.Context, token uuid.UUID) (ordered_set.OrderedSet[movie.ID], error) {
	var dbResult solutionResult
	if err := db.GetContext(ctx, &dbResult, "SELECT solution FROM applicants WHERE token=$1;", token); err != nil {
		return ordered_set.OrderedSet[movie.ID]{}, err
	}

	if !dbResult.Solution.Valid {
		return ordered_set.OrderedSet[movie.ID]{}, utilities.NotFound("solution")
	}

	var solution ordered_set.OrderedSet[movie.ID]
	if err := go_json.Unmarshal([]byte(dbResult.Solution.String), &solution); err != nil {
		return ordered_set.OrderedSet[movie.ID]{}, err
	}

	return solution, nil
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
