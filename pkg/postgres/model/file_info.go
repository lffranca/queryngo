package model

import (
	"database/sql"
	"github.com/lffranca/queryngo/domain"
)

type FileInfo struct {
	ID           sql.NullInt64
	Name         sql.NullString
	Extension    sql.NullString
	Key          sql.NullString
	Path         sql.NullString
	Size         sql.NullInt64
	ContentType  sql.NullString
	LastModified sql.NullTime
	Prefix       sql.NullString
	Bucket       sql.NullString
	Status       sql.NullString
}

func (item *FileInfo) Entity() *domain.FileInfo {
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

	return &domain.FileInfo{
		ID:           id,
		Name:         &item.Name.String,
		Extension:    &item.Extension.String,
		Key:          &item.Key.String,
		Path:         &item.Path.String,
		Size:         size,
		ContentType:  &item.ContentType.String,
		LastModified: &item.LastModified.Time,
		Prefix:       &item.Prefix.String,
		Bucket:       &item.Bucket.String,
		Status:       domain.FileStatus(item.Status.String),
	}
}
