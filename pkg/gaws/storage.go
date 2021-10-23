package gaws

import (
	"context"
	"io"
)

type StorageService service

func (pkg *StorageService) Upload(ctx context.Context, key, contentType *string, data io.Reader) error {
	return pkg.client.upload(ctx, key, contentType, data)
}
