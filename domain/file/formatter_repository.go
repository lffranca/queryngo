package file

import (
	"context"
	"io"
)

type ReaderRepository interface {
	Read(ctx context.Context, reader io.Reader) ([][]string, error)
}
