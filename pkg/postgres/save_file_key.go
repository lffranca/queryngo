package postgres

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type SaveFileKeyService service

func (pkg *SaveFileKeyService) SaveFileKey(ctx context.Context, data *domain.FileInfo) error {
	return pkg.client.File.Save(ctx, data)
}
