package querying

import (
	"context"
	"database/sql"
	"errors"
	"github.com/snowflakedb/gosnowflake"
	"log"
)

func NewSnowflake(db DB) (*queryingSnowflake, error) {
	if db == nil {
		return nil, errors.New("invalid params")
	}

	return &queryingSnowflake{
		db: db.DB(),
	}, nil
}

type queryingSnowflake struct {
	db *sql.DB
}

func (item *queryingSnowflake) query(ctx context.Context, query string, variables []interface{}) (*sql.Rows, error) {
	if variables == nil || len(variables) <= 0 {
		return item.db.QueryContext(ctx, query)
	}

	var args []interface{}
	for _, varItem := range variables {
		args = append(args, gosnowflake.Array(varItem))
	}

	return item.db.QueryContext(ctx, query, args...)
}

func (item *queryingSnowflake) Query(ctx context.Context, query string, variables []interface{}) ([]string, []string, [][]interface{}, error) {
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
