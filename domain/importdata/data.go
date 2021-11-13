package importdata

import (
	"context"
	"github.com/lffranca/queryngo/domain"
	"io"
)

type AbstractDatabase interface {
	Save(ctx context.Context, data *domain.FileInfo) error
}

type AbstractGenerate interface {
	UUID(ctx context.Context) (*string, error)
}

type AbstractStorage interface {
	Upload(ctx context.Context, key, contentType *string, data io.Reader) error
}
