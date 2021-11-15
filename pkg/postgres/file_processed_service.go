package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lffranca/queryngo/domain"
)

type FileProcessedService service

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
