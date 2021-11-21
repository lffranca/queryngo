package model

import (
	"database/sql"
	"encoding/json"
	"github.com/lffranca/queryngo/domain"
	"log"
)

type FileConfig struct {
	ID               sql.NullInt64
	FileProcessedID  sql.NullInt64
	RowOffset        sql.NullInt64
	ColumnDate       sql.NullString
	ColumnValue      sql.NullString
	ColumnsDimension sql.NullString
}

func (item *FileConfig) Entity() *domain.FileConfig {
	var id *int
	if item.ID.Valid {
		i := int(item.ID.Int64)
		id = &i
	}

	var processedID *int
	if item.FileProcessedID.Valid {
		i := int(item.FileProcessedID.Int64)
		processedID = &i
	}

	var rowOffset *int
	if item.RowOffset.Valid {
		i := int(item.RowOffset.Int64)
		rowOffset = &i
	}

	var columnsDimension []*string
	if item.ColumnsDimension.Valid {
		if err := json.Unmarshal([]byte(item.ColumnsDimension.String), &columnsDimension); err != nil {
			log.Println(err)
		}
	}

	return &domain.FileConfig{
		ID:               id,
		FileProcessedID:  processedID,
		RowOffset:        rowOffset,
		ColumnDate:       &item.ColumnDate.String,
		ColumnValue:      &item.ColumnValue.String,
		ColumnsDimension: columnsDimension,
	}
}
