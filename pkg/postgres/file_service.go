package postgres

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type FileService service

func (pkg *FileService) Save(ctx context.Context, data *domain.FileInfo) error {
	query := `
		insert into storage.files (key, path, name, extension, "size", content_type, last_modified)
		values ($1, $2, $3, $4, $5, $6, $7);
	`

	if _, err := pkg.client.db.ExecContext(ctx, query,
		data.Key,
		data.Path,
		data.Name,
		data.Extension,
		data.Size,
		data.ContentType,
		data.LastModified,
	); err != nil {
		return err
	}

	return nil
}
