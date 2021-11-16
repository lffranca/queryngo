package importdata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lffranca/queryngo/domain"
	"io"
	"path/filepath"
	"time"
)

func New(options *Options) (*ImportData, error) {
	if options == nil {
		return nil, errors.New("options param is required")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	return &ImportData{
		prefix:                options.Prefix,
		bucket:                options.Bucket,
		basePath:              options.BasePath,
		storageRepository:     options.StorageRepository,
		fileRepository:        options.FileRepository,
		uuidRepository:        options.UUIDRepository,
		processFileRepository: options.ProcessFileRepository,
	}, nil
}

type ImportData struct {
	prefix                *string
	bucket                *string
	basePath              *string
	storageRepository     StorageRepository
	fileRepository        FileRepository
	uuidRepository        UUIDRepository
	processFileRepository ProcessFileRepository
}

func (mod *ImportData) Import(ctx context.Context, fileName, contentType *string, fileSize *int, data io.Reader) error {
	extension := filepath.Ext(*fileName)

	uuid, err := mod.uuidRepository.UUID(ctx)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%s", *uuid, extension)
	path := fmt.Sprintf("%s/%s", *mod.basePath, key)

	if err := mod.storageRepository.Upload(ctx, &path, contentType, data); err != nil {
		return err
	}

	now := time.Now()
	fileInfo := &domain.FileInfo{
		Name:         fileName,
		Extension:    &extension,
		Key:          &key,
		Path:         &path,
		Size:         fileSize,
		ContentType:  contentType,
		LastModified: &now,
		Prefix:       mod.prefix,
		Bucket:       mod.bucket,
		Status:       domain.FileStatusPending,
	}

	fileInfo, err = mod.fileRepository.Save(ctx, fileInfo)
	if err != nil {
		return err
	}

	fileInfoJSON, err := json.Marshal(fileInfo)
	if err != nil {
		return err
	}

	if err := mod.processFileRepository.ProducerProcessFile(ctx, fileInfoJSON); err != nil {
		return err
	}

	return nil
}
