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
	client.repository = options.Repository
	client.processedRepository = options.ProcessedRepository
	client.configRepository = options.ConfigRepository
	client.storageRepository = options.StorageRepository
	client.readerRepository = options.ReaderRepository

	return client, nil
}

type File struct {
	repository          Repository
	processedRepository ProcessedRepository
	configRepository    ConfigRepository
	storageRepository   StorageRepository
	readerRepository    ReaderRepository
}

func (pkg *File) ListConfig(ctx context.Context, processedID *int, offset, limit *int, search *string) ([]*domain.FileConfig, error) {
	return pkg.configRepository.ListByProcessedID(ctx, processedID, offset, limit, search)
}

func (pkg *File) DeleteConfig(ctx context.Context, id *int) error {
	return pkg.configRepository.Delete(ctx, id)
}

func (pkg *File) SaveConfig(ctx context.Context, item *domain.FileConfig) (*domain.FileConfig, error) {
	return pkg.configRepository.Save(ctx, item)
}

func (pkg *File) ProcessedFileContent(ctx context.Context, id *int) ([][]string, error) {
	processedFile, err := pkg.processedRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	response, err := pkg.storageRepository.Download(ctx, processedFile.Path)
	if err != nil {
		return nil, err
	}

	data, err := pkg.readerRepository.Read(ctx, response)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (pkg *File) ListProcessedFile(ctx context.Context, parentID *int, offset, limit *int, search *string) ([]*domain.FileInfoResult, error) {
	return pkg.processedRepository.ListByParentID(ctx, parentID, offset, limit, search)
}

func (pkg *File) List(ctx context.Context, offset, limit *int, search *string) ([]*domain.FileInfo, error) {
	return pkg.repository.List(ctx, offset, limit, search)
}
