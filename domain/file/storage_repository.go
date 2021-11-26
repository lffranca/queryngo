package file

import (
	"context"
	"io"
)

type StorageRepository interface {
	Download(ctx context.Context, key *string) (io.Reader, error)
}
