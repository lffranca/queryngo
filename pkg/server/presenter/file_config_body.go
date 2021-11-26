package presenter

import "github.com/lffranca/queryngo/domain"

type FileConfigBody struct {
	FileProcessedID  *int      `form:"file_processed_id" json:"file_processed_id" binding:"required"`
	RowOffset        *int      `form:"row_offset" json:"row_offset" binding:"required"`
	ColumnDate       *string   `form:"column_date" json:"column_date" binding:"required"`
	ColumnValue      *string   `form:"column_value" json:"column_value" binding:"required"`
	ColumnsDimension []*string `form:"columns_dimension" json:"columns_dimension" binding:"required"`
}

func (item *FileConfigBody) Entity() *domain.FileConfig {
	return &domain.FileConfig{
		FileProcessedID:  item.FileProcessedID,
		RowOffset:        item.RowOffset,
		ColumnDate:       item.ColumnDate,
		ColumnValue:      item.ColumnValue,
		ColumnsDimension: item.ColumnsDimension,
	}
}
