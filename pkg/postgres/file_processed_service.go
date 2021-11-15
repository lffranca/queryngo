package postgres

import (
	"context"
	"encoding/json"
	"github.com/lffranca/queryngo/domain"
	"github.com/lib/pq"
)

type FileProcessedService service

func (pkg *FileProcessedService) SaveAll(ctx context.Context, items []*domain.FileInfoResult) error {
	var sheet []*string
	var path []*string
	var extension []*string
	var contentType []*string
	var columns []string
	var key []*string
	var size []*int
	var parentID []*int
	var bucket []*string
	var prefix []*string

	for _, item := range items {
		sheet = append(sheet, item.Sheet)
		path = append(path, item.Path)
		extension = append(extension, item.Extension)
		contentType = append(contentType, item.ContentType)

		columnsJSON, _ := json.Marshal(item.Columns)
		columns = append(columns, string(columnsJSON))

		key = append(key, item.Key)
		size = append(size, item.Size)
		parentID = append(parentID, item.ParentID)
		bucket = append(bucket, item.Bucket)
		prefix = append(prefix, item.Prefix)
	}

	query := `
		insert into storage.file_processed (
			sheet,
			path,
			extension,
			content_type,
			columns,
			key,
			size,
			file_id,
			bucket,
			prefix
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	if _, err := pkg.client.db.ExecContext(ctx, query,
		pq.Array(sheet),
		pq.Array(path),
		pq.Array(extension),
		pq.Array(contentType),
		pq.Array(columns),
		pq.Array(key),
		pq.Array(size),
		pq.Array(parentID),
		pq.Array(bucket),
		pq.Array(prefix),
	); err != nil {
		return err
	}

	return nil
}
