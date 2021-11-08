package server

import "context"

type QueryingRepository interface {
	Query(ctx context.Context, query string, variables []interface{}) (columns []string, columnTypes []string, values [][]interface{}, err error)
}
