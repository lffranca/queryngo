package gmigrate

import (
	"database/sql"
	"errors"
)

func New(db *sql.DB, path *string) (*Client, error) {
	if db == nil || path == nil {
		return nil, errors.New("invalid params")
	}

	client := new(Client)
	client.db = db
	client.path = path
	client.common.client = client
	client.Postgres = (*PostgresService)(&client.common)

	return client, nil
}

type service struct {
	client *Client
}

type Client struct {
	db       *sql.DB
	path     *string
	common   service
	Postgres *PostgresService
}
