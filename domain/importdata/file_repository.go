package importdata

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type FileRepository interface {
	Save(ctx context.Context, data *domain.FileInfo) (*domain.FileInfo, error)
}
