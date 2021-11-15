package server

import "context"

type QueryingRepository interface {
	Execute(ctx context.Context, queryID, formatID *string, value interface{}) ([]byte, error)
}
