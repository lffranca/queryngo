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

func New(storage AbstractStorage, db AbstractDatabase, generate AbstractGenerate) (*importData, error) {
	if storage == nil || db == nil || generate == nil {
		return nil, errors.New("invalid params")
	}

	return &importData{
		storage:  storage,
		db:       db,
		generate: generate,
	}, nil
}

type importData struct {
	storage  AbstractStorage
	db       AbstractDatabase
	generate AbstractGenerate
}

func (mod *importData) Import(ctx context.Context, fileName, contentType *string, fileSize *int, data io.Reader) error {
	extension := filepath.Ext(*fileName)

	uuid, err := mod.generate.UUID(ctx)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%s", *uuid, extension)

	if err := mod.storage.Upload(ctx, &key, contentType, data); err != nil {
		return err
	}

	now := time.Now()
	fileInfo := &domain.FileInfo{
		Name:         fileName,
		Extension:    &extension,
		Key:          &key,
		Path:         &key,
		Size:         fileSize,
		ContentType:  contentType,
		LastModified: &now,
	}

	if err := mod.db.Save(ctx, fileInfo); err != nil {
		return err
	}

	return nil
}
