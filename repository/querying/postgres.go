package querying

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"log"
)

func NewPostgres(db DB) (*queryingPostgres, error) {
	if db == nil {
		return nil, errors.New("invalid params")
	}

	return &queryingPostgres{
		db: db.DB(),
	}, nil
}

type queryingPostgres struct {
	db *sql.DB
}

func (item *queryingPostgres) query(ctx context.Context, query string, variables []interface{}) (*sql.Rows, error) {
	if variables == nil || len(variables) <= 0 {
		return item.db.QueryContext(ctx, query)
	}

	var args []interface{}
	for _, varItem := range variables {
		args = append(args, pq.Array(varItem))
	}

	return item.db.QueryContext(ctx, query, args...)
}

func (item *queryingPostgres) Query(ctx context.Context, query string, variables []interface{}) ([]string, []string, [][]interface{}, error) {
	rows, errRows := item.query(ctx, query, variables)
	if errRows != nil {
		return nil, nil, nil, errRows
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("ERROR CLOSE ROWS: ", err)
		}
	}()

	return NewDataModel(rows)
}
