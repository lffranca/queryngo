package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lffranca/queryngo/domain"
	"github.com/lffranca/queryngo/pkg/postgres/model"
	"log"
)

type FileProcessedService service

func (pkg *FileProcessedService) Get(ctx context.Context, id *int) (*domain.FileInfoResult, error) {
	if id == nil {
		return nil, errors.New("id param is required")
	}

	query := `
		select
			id,
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
		from storage.file_processed
		where id = $1;
	`

	var itemDB model.FileInfoResult
	if err := pkg.client.db.QueryRowContext(ctx, query, *id).Scan(
		&itemDB.ID,
		&itemDB.Sheet,
		&itemDB.Path,
		&itemDB.Extension,
		&itemDB.ContentType,
		&itemDB.Columns,
		&itemDB.Key,
		&itemDB.Size,
		&itemDB.ParentID,
		&itemDB.Bucket,
		&itemDB.Prefix,
	); err != nil {
		return nil, err
	}

	return itemDB.Entity(), nil
}

func (pkg *FileProcessedService) Delete(ctx context.Context, id *int) error {
	if id == nil {
		return errors.New("id param is required")
	}

	query := "delete from storage.file_processed where id = $1;"

	if _, err := pkg.client.db.ExecContext(ctx, query, *id); err != nil {
		return err
	}

	return nil
}

func (pkg *FileProcessedService) List(ctx context.Context, offset, limit *int, search *string) ([]*domain.FileInfoResult, error) {
	if offset == nil {
		offset = &defaultOffset
	}

	if limit == nil {
		limit = &defaultLimit
	}

	query := `
		select
		    id,
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
		from storage.file_processed
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

	var items []*domain.FileInfoResult
	if rows.Next() {
		var itemDB model.FileInfoResult
		if err := rows.Scan(
			&itemDB.ID,
			&itemDB.Sheet,
			&itemDB.Path,
			&itemDB.Extension,
			&itemDB.ContentType,
			&itemDB.Columns,
			&itemDB.Key,
			&itemDB.Size,
			&itemDB.ParentID,
			&itemDB.Bucket,
			&itemDB.Prefix,
		); err != nil {
			return nil, err
		}

		items = append(items, itemDB.Entity())
	}

	return items, nil
}

func (pkg *FileProcessedService) ListByParentID(ctx context.Context, parentID *int, offset, limit *int, search *string) ([]*domain.FileInfoResult, error) {
	if parentID == nil {
		return nil, errors.New("parent id param is required")
	}

	if offset == nil {
		offset = &defaultOffset
	}

	if limit == nil {
		limit = &defaultLimit
	}

	query := `
		select
		    id,
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
		from storage.file_processed
		where file_id = $1
	`

	args := []interface{}{*parentID}
	argsCount := 1

	if search != nil {
		argsCount++
		query += fmt.Sprintf(" and name ilike %d ", argsCount)
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

	var items []*domain.FileInfoResult
	if rows.Next() {
		var itemDB model.FileInfoResult
		if err := rows.Scan(
			&itemDB.ID,
			&itemDB.Sheet,
			&itemDB.Path,
			&itemDB.Extension,
			&itemDB.ContentType,
			&itemDB.Columns,
			&itemDB.Key,
			&itemDB.Size,
			&itemDB.ParentID,
			&itemDB.Bucket,
			&itemDB.Prefix,
		); err != nil {
			return nil, err
		}

		items = append(items, itemDB.Entity())
	}

	return items, nil
}

func (pkg *FileProcessedService) Update(ctx context.Context, data *domain.FileInfoResult) error {
	if data == nil {
		return errors.New("item is required")
	}

	if data.ID == nil {
		return errors.New("item id is required")
	}

	query := `
		update storage.file_processed set
			sheet = $1,
			path = $2,
			extension = $3,
			content_type = $4,
			columns = $5,
			key = $6,
			size = $7,
			file_id = $8,
			bucket = $9,
			prefix = $10
		where id = $11;
	`

	columnsJSON, err := json.Marshal(data.Columns)
	if err != nil {
		return err
	}

	if _, err := pkg.client.db.ExecContext(ctx, query,
		data.Sheet,
		data.Path,
		data.Extension,
		data.ContentType,
		string(columnsJSON),
		data.Key,
		data.Size,
		data.ParentID,
		data.Bucket,
		data.Prefix,
	); err != nil {
		return err
	}

	return nil
}

func (pkg *FileProcessedService) SaveAll(ctx context.Context, items []*domain.FileInfoResult) error {
	if len(items) <= 0 {
		return nil
	}

	query, args := pkg.getQueryAndArgs(items)

	if _, err := pkg.client.db.ExecContext(context.Background(), *query, args...); err != nil {
		return err
	}

	return nil
}

func (pkg *FileProcessedService) getQueryAndArgs(items []*domain.FileInfoResult) (*string, []interface{}) {
	queryAll := `
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
		) values 
	`

	var args []interface{}
	for index, i := range items {

		query := " "

		query, args = pkg.appendArgs("($%d, ", query, args, i.Sheet)
		query, args = pkg.appendArgs("$%d, ", query, args, i.Path)
		query, args = pkg.appendArgs("$%d, ", query, args, i.Extension)
		query, args = pkg.appendArgs("$%d, ", query, args, i.ContentType)

		columnsJSON, _ := json.Marshal(i.Columns)
		query, args = pkg.appendArgs("$%d, ", query, args, string(columnsJSON))

		query, args = pkg.appendArgs("$%d, ", query, args, i.Key)
		query, args = pkg.appendArgs("$%d, ", query, args, i.Size)
		query, args = pkg.appendArgs("$%d, ", query, args, i.ParentID)
		query, args = pkg.appendArgs("$%d, ", query, args, i.Bucket)
		query, args = pkg.appendArgs("$%d)", query, args, i.Prefix)

		if (index + 1) < len(items) {
			query += `,
			`
		}

		queryAll += query
	}

	queryAll += ";"

	return &queryAll, args
}

func (pkg *FileProcessedService) appendArgs(tmpl, query string, args []interface{}, value interface{}) (string, []interface{}) {
	args = append(args, value)
	query += fmt.Sprintf(tmpl, len(args))
	return query, args
}
