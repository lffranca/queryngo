package server

import "context"

type TemplateRepository interface {
	ByID(ctx context.Context, id *string) ([]byte, error)
}
