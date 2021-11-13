package server

import "context"

type UUIDRepository interface {
	UUID(ctx context.Context) (*string, error)
}
