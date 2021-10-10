package snowflake

import (
	"database/sql"
	"errors"

	_ "github.com/snowflakedb/gosnowflake"
)

func New(conn *string) (*snowflake, error) {
	if conn == nil {
		return nil, errors.New("conn is required")
	}

	db, err := sql.Open("snowflake", *conn)
	if err != nil {
		return nil, err
	}

	return &snowflake{
		db: db,
	}, nil
}

type snowflake struct {
	db *sql.DB
}

func (pkg *snowflake) DB() *sql.DB {
	return pkg.db
}

func (pkg *snowflake) Close() {
	if err := pkg.db.Close(); err != nil {
		panic(err)
	}
}
