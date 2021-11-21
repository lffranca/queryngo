package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lffranca/queryngo/domain"
	"github.com/lffranca/queryngo/pkg/postgres/model"
	"log"
)

type FileService service

func (pkg *FileService) Delete(ctx context.Context, id *int) error {
	if id == nil {
		return errors.New("id param is required")
	}

	query := "delete from storage.files where id = $1;"

	if _, err := pkg.client.db.ExecContext(ctx, query, *id); err != nil {
		return err
	}

	return nil
}

func (pkg *FileService) Get(ctx context.Context, id *int) (*domain.FileInfo, error) {
	if id == nil {
		return nil, errors.New("id param is required")
	}

	query := `
		select
			id,
		    key,
		    path,
		    name,
		    extension,
		    size,
		    content_type,
		    last_modified,
		    prefix,
		    bucket,
		    status
		from storage.files
		where id = $1;
	`

	var itemDB model.FileInfo
	if err := pkg.client.db.QueryRowContext(ctx, query, *id).Scan(
		&itemDB.ID,
		&itemDB.Key,
		&itemDB.Path,
		&itemDB.Name,
		&itemDB.Extension,
		&itemDB.Size,
		&itemDB.ContentType,
		&itemDB.LastModified,
		&itemDB.Prefix,
		&itemDB.Bucket,
		&itemDB.Status,
	); err != nil {
		return nil, err
	}

	return itemDB.Entity(), nil
}

func (pkg *FileService) List(ctx context.Context, offset, limit *int, search *string) ([]*domain.FileInfo, error) {
	if offset == nil {
		offset = &defaultOffset
	}

	if limit == nil {
		limit = &defaultLimit
	}

	query := `
		select
			id,
		    key,
		    path,
		    name,
		    extension,
		    size,
		    content_type,
		    last_modified,
		    prefix,
		    bucket,
		    status
		from storage.files
	`

	var args []interface{}
	var argsCount int

	if search != nil {
		argsCount++
		query += fmt.Sprintf(" where name ilike %d ", argsCount)
		args = append(args, *search)
	}

	if limit != nil {
		argsCount++
		query += fmt.Sprintf(" limit %d ", argsCount)
		args = append(args, *limit)
	}

	if offset != nil {
		argsCount++
		query += fmt.Sprintf(" offset %d ", argsCount)
		args = append(args, *offset)
	}

	rows, err := pkg.client.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var items []*domain.FileInfo
	if rows.Next() {
		var itemDB model.FileInfo
		if err := rows.Scan(
			&itemDB.ID,
			&itemDB.Key,
			&itemDB.Path,
			&itemDB.Name,
			&itemDB.Extension,
			&itemDB.Size,
			&itemDB.ContentType,
			&itemDB.LastModified,
			&itemDB.Prefix,
			&itemDB.Bucket,
			&itemDB.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, itemDB.Entity())
	}

	return items, nil
}

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

func (pkg *FileService) Save(ctx context.Context, data *domain.FileInfo) (*domain.FileInfo, error) {
	if data == nil {
		return nil, errors.New("data is required")
	}

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
	   	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	   	returning id;
	`

	var idSQL sql.NullInt64
	if err := pkg.client.db.QueryRowContext(ctx, query,
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
	).Scan(&idSQL); err != nil {
		return nil, err
	}

	id := int(idSQL.Int64)
	data.ID = &id

	return data, nil
}
