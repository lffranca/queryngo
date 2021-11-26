package file

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type Repository interface {
	List(ctx context.Context, offset, limit *int, search *string) ([]*domain.FileInfo, error)
}
