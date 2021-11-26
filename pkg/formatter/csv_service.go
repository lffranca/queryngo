package formatter

import (
	"context"
	"encoding/csv"
	"io"
)

type CSVService service

func (pkg *CSVService) Read(ctx context.Context, reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
