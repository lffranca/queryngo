package importdata

import (
	"context"
	"io"
)

type StorageRepository interface {
	Upload(ctx context.Context, key, contentType *string, data io.Reader) error
}
