package postgres

import (
	"context"
	"errors"
	"github.com/lffranca/queryngo/domain"
)

type FileService service

func (pkg *FileService) Update(ctx context.Context, data *domain.FileInfo) error {
	if data == nil {
		return errors.New("item is required")
	}

	if data.ID == nil {
		return errors.New("item id is required")
	}

	query := `
		update storage.files set
			key = $1,
			path = $2,
			name = $3,
			extension = $4,
			"size" = $5,
			content_type = $6,
			last_modified = $7,
			prefix = $8,
			bucket = $9,
			status = $10
		where id = $11;
	`

	if _, err := pkg.client.db.ExecContext(ctx, query,
		data.Key,
		data.Path,
		data.Name,
		data.Extension,
		data.Size,
		data.ContentType,
		data.LastModified,
		data.Prefix,
		data.Bucket,
		data.Status,
		data.ID,
	); err != nil {
		return err
	}

	return nil
}

func (pkg *FileService) Save(ctx context.Context, data *domain.FileInfo) error {
	query := `
		insert into storage.files (
		   key,
		   path,
		   name,
		   extension,
		   "size",
		   content_type,
		   last_modified,
		   prefix,
		   bucket,
		   status
	   	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`

	if _, err := pkg.client.db.ExecContext(ctx, query,
		data.Key,
		data.Path,
		data.Name,
		data.Extension,
		data.Size,
		data.ContentType,
		data.LastModified,
		data.Prefix,
		data.Bucket,
		data.Status,
	); err != nil {
		return err
	}

	return nil
}
