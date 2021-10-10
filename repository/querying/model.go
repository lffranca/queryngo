package querying

import (
	"database/sql"
	"errors"
	"time"
)

func NewDataModel(rows *sql.Rows) ([]string, []string, [][]interface{}, error) {
	if rows == nil {
		return nil, nil, nil, errors.New("invalid param")
	}

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return nil, nil, nil, errColumns
	}

	columnTypes, errColumnTypes := rows.ColumnTypes()
	if errColumnTypes != nil {
		return nil, nil, nil, errColumnTypes
	}

	var columnTypesString []string
	for _, col := range columnTypes {
		columnTypesString = append(columnTypesString, col.ScanType().Name())
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var items [][]interface{}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, nil, nil, err
		}

		var vi []interface{}
		for _, v := range values {
			switch t := v.(type) {
			case string:
				vi = append(vi, (string)(t))
				continue
			case int:
				vi = append(vi, (int)(t))
				continue
			case int8:
				vi = append(vi, (int8)(t))
				continue
			case int16:
				vi = append(vi, (int16)(t))
				continue
			case int32:
				vi = append(vi, (int32)(t))
				continue
			case int64:
				vi = append(vi, (int64)(t))
				continue
			case float32:
				vi = append(vi, (float32)(t))
				continue
			case float64:
				vi = append(vi, (float64)(t))
				continue
			case bool:
				vi = append(vi, (bool)(t))
				continue
			case time.Time:
				vi = append(vi, (time.Time)(t))
				continue
			default:
				vi = append(vi, nil)
			}
		}

		items = append(items, vi)
	}

	return columns, columnTypesString, items, nil
}
