package file

import (
	"context"
	"errors"
	"github.com/lffranca/queryngo/domain"
)

func New(options *Options) (*File, error) {
	if options == nil {
		return nil, errors.New("options is required")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	client := new(File)
	return client, nil
}

type File struct{}

func (pkg *File) ListConfig(ctx context.Context, processedID *int, offset, limit *int, search *string) ([]*domain.FileConfig, error) {
	return nil, nil
}

func (pkg *File) DeleteConfig(ctx context.Context, id *int) error {
	return nil
}

func (pkg *File) SaveConfig(ctx context.Context, item *domain.FileConfig) (*domain.FileConfig, error) {
	return nil, nil
}

func (pkg *File) ProcessedFileContent(ctx context.Context, id *int) ([]byte, error) {
	return nil, nil
}

func (pkg *File) ListProcessedFile(ctx context.Context, parentID *int, offset, limit *int, search *string) ([]*domain.FileInfoResult, error) {
	return nil, nil
}

func (pkg *File) List(ctx context.Context, offset, limit *int, search *string) ([]*domain.FileInfo, error) {
	return nil, nil
}
