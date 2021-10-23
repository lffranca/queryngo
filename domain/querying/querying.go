package querying

import (
	"context"
	"errors"
)

func New(format Format, formatter Formatter, query AbstractQuerying) (*querying, error) {
	if format == nil || formatter == nil || query == nil {
		return nil, errors.New("invalid params")
	}

	return &querying{
		format:    format,
		formatter: formatter,
		querying:  query,
	}, nil
}

type querying struct {
	format    Format
	formatter Formatter
	querying  AbstractQuerying
}

func (mod *querying) Execute(ctx context.Context, queryID, formatID *string, value interface{}) ([]byte, error) {
	queryTemplate, err := mod.format.ByID(ctx, queryID)
	if err != nil {
		return nil, err
	}

	formatTemplate, err := mod.format.ByID(ctx, formatID)
	if err != nil {
		return nil, err
	}

	query, err := mod.formatter.Transform(ctx, queryTemplate, value)
	if err != nil {
		return nil, err
	}

	columns, columnTypes, values, err := mod.querying.Query(ctx, string(query), nil)
	if err != nil {
		return nil, err
	}

	return mod.formatter.Transform(ctx, formatTemplate, map[string]interface{}{
		"Columns":     columns,
		"ColumnTypes": columnTypes,
		"Values":      values,
	})
}
