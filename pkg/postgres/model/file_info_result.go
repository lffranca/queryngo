package model

import (
	"database/sql"
	"encoding/json"
	"github.com/lffranca/queryngo/domain"
	"log"
)

type FileInfoResult struct {
	ID          sql.NullInt64
	Sheet       sql.NullString
	Path        sql.NullString
	Extension   sql.NullString
	ContentType sql.NullString
	Columns     sql.NullString
	Key         sql.NullString
	Size        sql.NullInt64
	ParentID    sql.NullInt64
	Bucket      sql.NullString
	Prefix      sql.NullString
}

func (item *FileInfoResult) Entity() *domain.FileInfoResult {
	var id *int
	if item.ID.Valid {
		i := int(item.ID.Int64)
		id = &i
	}

	var size *int
	if item.Size.Valid {
		i := int(item.Size.Int64)
		size = &i
	}

	var columns []*string
	if item.Columns.Valid {
		if err := json.Unmarshal([]byte(item.Columns.String), &columns); err != nil {
			log.Println(err)
		}
	}

	var parentID *int
	if item.ParentID.Valid {
		i := int(item.ParentID.Int64)
		parentID = &i
	}

	return &domain.FileInfoResult{
		ID:          id,
		Sheet:       &item.Sheet.String,
		Path:        &item.Path.String,
		Extension:   &item.Extension.String,
		ContentType: &item.ContentType.String,
		Columns:     columns,
		Key:         &item.Key.String,
		Size:        size,
		ParentID:    parentID,
		Bucket:      &item.Bucket.String,
		Prefix:      &item.Prefix.String,
	}
}
