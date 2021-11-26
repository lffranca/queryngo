package server

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type StorageRepository interface {
	ListConfig(ctx context.Context, processedID *int, offset, limit *int, search *string) ([]*domain.FileConfig, error)
	DeleteConfig(ctx context.Context, id *int) error
	SaveConfig(ctx context.Context, item *domain.FileConfig) (*domain.FileConfig, error)
	ProcessedFileContent(ctx context.Context, id *int) ([][]string, error)
	ListProcessedFile(ctx context.Context, parentID *int, offset, limit *int, search *string) ([]*domain.FileInfoResult, error)
	List(ctx context.Context, offset, limit *int, search *string) ([]*domain.FileInfo, error)
}
