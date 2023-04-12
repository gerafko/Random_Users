package storage

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Postgres represents PostgreSQL connection.
type Postgres struct {
	*sqlx.DB
}

// NewPostgres connects to PostgreSQL database and returns its connection.
func NewPostgres(ctx context.Context, dsn string, delay time.Duration, retry int) (*Postgres, error) {
	var err error

	for i := 0; i < retry; i++ {
		db, err := sqlx.ConnectContext(ctx, "pgx", dsn)
		if err != nil {
			time.Sleep(delay)
			continue
		}

		return &Postgres{
			DB: db,
		}, nil
	}

	return nil, errors.Wrap(err, "couldn't connect to PostgreSQL database")
}
