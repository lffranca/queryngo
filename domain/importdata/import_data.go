package importdata

import (
	"context"
	"errors"
	"fmt"
	"github.com/lffranca/queryngo/domain"
	"io"
	"path/filepath"
	"time"
)

func New(storage AbstractStorage, db AbstractDatabase, generate AbstractGenerate, basePath *string) (*importData, error) {
	if storage == nil || db == nil || generate == nil || basePath == nil {
		return nil, errors.New("invalid params")
	}

	return &importData{
		storage:  storage,
		db:       db,
		generate: generate,
		basePath: basePath,
	}, nil
}

type importData struct {
	storage  AbstractStorage
	db       AbstractDatabase
	generate AbstractGenerate
	basePath *string
}

func (mod *importData) Import(ctx context.Context, fileName, contentType *string, fileSize *int, data io.Reader) error {
	extension := filepath.Ext(*fileName)

	uuid, err := mod.generate.UUID(ctx)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%s", *uuid, extension)

	path := fmt.Sprintf("%s/%s", *mod.basePath, key)

	if err := mod.storage.Upload(ctx, &path, contentType, data); err != nil {
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
	}

	if err := mod.db.SaveFileKey(ctx, fileInfo); err != nil {
		return err
	}

	return nil
}
