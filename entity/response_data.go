package entity

func NewResponseData(columns, columnTypes []string, rows [][]interface{}) *ResponseData {
	return &ResponseData{
		Columns:     columns,
		ColumnTypes: columnTypes,
		Rows:        rows,
	}
}

type ResponseData struct {
	Columns     []string
	ColumnTypes []string
	Rows        [][]interface{}
}
