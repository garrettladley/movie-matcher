package storage

import (
	"movie-matcher/internal/config"

	"github.com/jmoiron/sqlx"
)

func CreatePostgresConnection(settings config.Settings) *sqlx.DB {
	return sqlx.MustConnect("postgres", settings.Database.WithDb())
}
