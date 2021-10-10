package querying

import "context"

type Querying interface {
	Query(ctx context.Context, query string, variables []interface{}) (columns []string, columnTypes []string, values [][]interface{}, err error)
}
