package domain

type FileConfig struct {
	ID               *int      `json:"id,omitempty" yaml:"id,omitempty"`
	FileProcessedID  *int      `json:"file_processed_id,omitempty" yaml:"file_processed_id,omitempty"`
	RowOffset        *int      `json:"row_offset,omitempty" yaml:"row_offset,omitempty"`
	ColumnDate       *string   `json:"column_date,omitempty" yaml:"column_date,omitempty"`
	ColumnValue      *string   `json:"column_value,omitempty" yaml:"column_value,omitempty"`
	ColumnsDimension []*string `json:"columns_dimension,omitempty" yaml:"columns_dimension,omitempty"`
}
