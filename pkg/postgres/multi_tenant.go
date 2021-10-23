package postgres

import (
	"context"
	"database/sql"
	"errors"
)

func NewMultiTenant(db *sql.DB) (*multiTenant, error) {
	if db == nil {
		return nil, errors.New("invalid params")
	}

	return &multiTenant{
		db:  db,
		dbs: make(map[string]*Client),
	}, nil
}

type multiTenant struct {
	db  *sql.DB
	dbs map[string]*Client
}

func (post *multiTenant) Client(ctx context.Context, sub *string) (*Client, error) {
	db, ok := post.dbs[*sub]
	if ok {
		return db, nil
	}

	db, err := post.getDB(ctx, sub)
	if err != nil {
		return nil, err
	}

	post.dbs[*sub] = db

	return db, nil
}

func (post *multiTenant) getDB(ctx context.Context, sub *string) (*Client, error) {
	query := `
		select
			d.host,
			d.port,
			d.username,
			d.password,
			d.database,
			d.encrypt
		from public.db_analytics as d
		inner join public.profile p on d.company_id = p.company_id
		where p.cognito_sub = $1
		limit 1
		;
	`

	var conn connectionDB
	if err := post.db.QueryRowContext(ctx, query, *sub).Scan(
		&conn.Host,
		&conn.Port,
		&conn.Username,
		&conn.Password,
		&conn.Database,
		&conn.Encrypt,
	); err != nil {
		return nil, err
	}

	if !conn.Host.Valid {
		return nil, errors.New("invalid user")
	}

	return New(conn.String())
}
