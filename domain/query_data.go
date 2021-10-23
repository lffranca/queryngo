package domain

import (
	"errors"
	"fmt"
	"strconv"
)

func NewQueryData(query string, columns []ColumnType, rows [][]interface{}) *QueryData {
	return &QueryData{
		Query:   query,
		Columns: columns,
		Rows:    rows,
	}
}

type QueryData struct {
	Query   string
	Columns []ColumnType
	Rows    [][]interface{}
}

func (qd *QueryData) RowsToColumns() ([]interface{}, error) {
	columns := make(map[int][]interface{})
	for _, row := range qd.Rows {
		for index, cell := range row {
			col, ok := columns[index]
			if ok {
				columns[index] = append(col, cell)
				continue
			}

			columns[index] = []interface{}{cell}
		}
	}

	variables := []interface{}{}
	for columnTypeID, data := range columns {
		switch qd.Columns[columnTypeID] {
		case ColumnTypeInt:
			variables = append(variables, qd.toArrayInt(data))
			continue
		case ColumnTypeFloat:
			variables = append(variables, qd.toArrayFloat(data))
			continue
		case ColumnTypeString:
			variables = append(variables, qd.toArrayString(data))
			continue
		default:
			return nil, errors.New("invalid column type")
		}
	}

	return variables, nil
}

func (qd *QueryData) toArrayString(data []interface{}) []string {
	items := []string{}
	for _, item := range data {
		items = append(items, fmt.Sprint(item))
	}

	return items
}

func (qd *QueryData) toArrayFloat(data []interface{}) []float64 {
	items := []float64{}
	for _, item := range data {
		switch t := item.(type) {
		case int:
			items = append(items, float64((int)(t)))
			continue
		case int32:
			items = append(items, float64((int32)(t)))
			continue
		case int64:
			items = append(items, float64((int64)(t)))
			continue
		case float64:
			items = append(items, (float64)(t))
			continue
		case float32:
			items = append(items, float64((float32)(t)))
			continue
		case string:
			f, _ := strconv.ParseFloat((string)(t), 64)
			items = append(items, f)
			continue
		case *int:
			i := (*int)(t)
			items = append(items, float64(*i))
			continue
		case *int32:
			i32 := (*int32)(t)
			i := float64(*i32)
			items = append(items, i)
			continue
		case *int64:
			i64 := (*int64)(t)
			i := float64(*i64)
			items = append(items, i)
			continue
		case *float64:
			f64 := (*float64)(t)
			i := float64(*f64)
			items = append(items, i)
			continue
		case *float32:
			f32 := (*float32)(t)
			i := float64(*f32)
			items = append(items, i)
			continue
		case *string:
			s := (*string)(t)
			f, _ := strconv.ParseFloat(*s, 64)
			items = append(items, f)
			continue
		default:
			items = append(items, 0)
		}
	}

	return items
}

func (qd *QueryData) toArrayInt(data []interface{}) []int {
	items := []int{}
	for _, item := range data {
		switch t := item.(type) {
		case int:
			items = append(items, (int)(t))
			continue
		case int32:
			items = append(items, int((int32)(t)))
			continue
		case int64:
			items = append(items, int((int64)(t)))
			continue
		case float64:
			items = append(items, int((float64)(t)))
			continue
		case float32:
			items = append(items, int((float32)(t)))
			continue
		case string:
			i, _ := strconv.ParseInt((string)(t), 10, 64)
			items = append(items, int(i))
			continue
		case *int:
			items = append(items, *((*int)(t)))
			continue
		case *int32:
			i32 := (*int32)(t)
			i := int(*i32)
			items = append(items, i)
			continue
		case *int64:
			i64 := (*int64)(t)
			i := int(*i64)
			items = append(items, i)
			continue
		case *float64:
			f64 := (*float64)(t)
			i := int(*f64)
			items = append(items, i)
			continue
		case *float32:
			f32 := (*float32)(t)
			i := int(*f32)
			items = append(items, i)
			continue
		case *string:
			s := (*string)(t)
			i, _ := strconv.ParseInt(*s, 10, 64)
			items = append(items, int(i))
			continue
		default:
			items = append(items, 0)
		}
	}

	return items
}
