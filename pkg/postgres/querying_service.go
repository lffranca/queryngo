package postgres

import (
	"context"
)

type QueryingService service

func (pkg *QueryingService) Query(ctx context.Context, query string, variables []interface{}) ([]string, []string, [][]interface{}, error) {
	return pkg.client.querying(ctx, query, variables)
}
