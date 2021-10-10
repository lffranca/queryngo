package querying

import "context"

type Format interface {
	ByID(ctx context.Context, id *string) ([]byte, error)
}
