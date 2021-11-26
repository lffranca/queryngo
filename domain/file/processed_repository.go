package file

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type ProcessedRepository interface {
	Get(ctx context.Context, id *int) (*domain.FileInfoResult, error)
	ListByParentID(ctx context.Context, parentID *int, offset, limit *int, search *string) ([]*domain.FileInfoResult, error)
}
