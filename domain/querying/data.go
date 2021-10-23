package querying

import "context"

type AbstractQuerying interface {
	Query(ctx context.Context, query string, variables []interface{}) (columns []string, columnTypes []string, values [][]interface{}, err error)
}

type Formatter interface {
	Transform(ctx context.Context, template []byte, input interface{}) (expected []byte, err error)
}

type Format interface {
	ByID(ctx context.Context, id *string) ([]byte, error)
}
