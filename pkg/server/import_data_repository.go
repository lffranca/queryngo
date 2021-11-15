package server

import (
	"context"
	"io"
)

type ImportDataRepository interface {
	Import(ctx context.Context, fileName, contentType *string, fileSize *int, data io.Reader) error
}
