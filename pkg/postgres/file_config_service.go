package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lffranca/queryngo/domain"
	"github.com/lffranca/queryngo/pkg/postgres/model"
	"log"
)

type FileConfigService service

func (pkg *FileConfigService) Delete(ctx context.Context, id *int) error {
	if id == nil {
		return errors.New("id param is required")
	}

	query := "delete from storage.file_config where id = $1;"

	if _, err := pkg.client.db.ExecContext(ctx, query, *id); err != nil {
		return err
	}

	return nil
}

func (pkg *FileConfigService) ListByProcessedID(ctx context.Context, processedID *int, offset, limit *int, search *string) ([]*domain.FileConfig, error) {
	if processedID == nil {
		return nil, errors.New("processed id param is required")
	}

	if offset == nil {
		offset = &defaultOffset
	}

	if limit == nil {
		limit = &defaultLimit
	}

	query := `
		select
		    id, file_processed_id, row_offset, column_date, column_value, columns_dimension
		from storage.file_config
		where file_processed_id = $1
	`

	args := []interface{}{*processedID}
	argsCount := 1

	if search != nil {
		argsCount++
		query += fmt.Sprintf(" and name ilike $%d ", argsCount)
		args = append(args, *search)
	}

	if limit != nil {
		argsCount++
		query += fmt.Sprintf(" limit $%d ", argsCount)
		args = append(args, *limit)
	}

	if offset != nil {
		argsCount++
		query += fmt.Sprintf(" offset $%d ", argsCount)
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

	var items []*domain.FileConfig
	for rows.Next() {
		var itemDB model.FileConfig
		if err := rows.Scan(
			&itemDB.ID,
			&itemDB.FileProcessedID,
			&itemDB.RowOffset,
			&itemDB.ColumnDate,
			&itemDB.ColumnValue,
			&itemDB.ColumnsDimension,
		); err != nil {
			return nil, err
		}

		items = append(items, itemDB.Entity())
	}

	return items, nil
}

func (pkg *FileConfigService) Save(ctx context.Context, item *domain.FileConfig) (*domain.FileConfig, error) {
	if item == nil {
		return nil, errors.New("item param is required")
	}

	query := `
		insert into storage.file_config (file_processed_id, row_offset, column_date, column_value, columns_dimension)
		values ($1, $2, $3, $4, $5);
	`

	var columns string
	if len(item.ColumnsDimension) > 0 {
		columnsResult, err := json.Marshal(item.ColumnsDimension)
		if err != nil {
			return nil, err
		}

		columns = string(columnsResult)
	}

	var idSQL sql.NullInt64
	if err := pkg.client.db.QueryRowContext(ctx, query,
		item.FileProcessedID,
		item.RowOffset,
		item.ColumnDate,
		item.ColumnValue,
		columns,
	).Scan(&idSQL); err != nil {
		return nil, err
	}

	if idSQL.Valid {
		i := int(idSQL.Int64)
		item.ID = &i
	}

	return item, nil
}
