package gkafka

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type ProcessedFileRepository interface {
	ProcessedFileResult(ctx context.Context, info *domain.FileInfo) error
}
