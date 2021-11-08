package postgres

import (
	"context"
	"database/sql"
	"errors"
)

type TemplateService service

func (pkg *TemplateService) ByID(ctx context.Context, id *string) ([]byte, error) {
	query := `
		select
			id, data
		from formatter.templates
		where id = $1
		limit 1
		;
	`

	item := struct {
		ID   sql.NullInt64
		Data sql.NullString
	}{}

	if err := pkg.client.db.QueryRowContext(ctx, query, *id).Scan(
		&item.ID,
		&item.Data,
	); err != nil {
		return nil, err
	}

	if !item.ID.Valid {
		return nil, errors.New("invalid template id")
	}

	return []byte(item.Data.String), nil
}
