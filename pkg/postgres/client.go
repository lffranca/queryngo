package postgres

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

func New(conn *string) (*postgres, error) {
	if conn == nil {
		return nil, errors.New("conn is required")
	}

	db, err := sql.Open("postgres", *conn)
	if err != nil {
		return nil, err
	}

	return &postgres{
		db: db,
	}, nil
}

type postgres struct {
	db *sql.DB
}

func (pkg *postgres) DB() *sql.DB {
	return pkg.db
}

func (pkg *postgres) Close() {
	if err := pkg.db.Close(); err != nil {
		panic(err)
	}
}
