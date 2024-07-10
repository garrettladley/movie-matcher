package storage

import (
	"movie-matcher/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreatePostgresConnection(settings config.DatabaseSettings) *sqlx.DB {
	return sqlx.MustConnect("postgres", settings.WithDb())
}
