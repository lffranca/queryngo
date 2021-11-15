package processedfile

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type FileProcessedRepository interface {
	SaveAll(ctx context.Context, items []*domain.FileInfoResult) error
}
