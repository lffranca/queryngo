package file

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type ConfigRepository interface {
	Delete(ctx context.Context, id *int) error
	ListByProcessedID(ctx context.Context, processedID *int, offset, limit *int, search *string) ([]*domain.FileConfig, error)
	Save(ctx context.Context, item *domain.FileConfig) (*domain.FileConfig, error)
}
