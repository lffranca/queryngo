package queryingo

import (
	"github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func serverQueryingRoute(route config.QueryingRoute, db *postgres.Client) (*querying.Querying, error) {
	if route.Enabled {
		formatterClient, err := formatter.New()
		if err != nil {
			return nil, err
		}

		mod, err := querying.New(db.Template, formatterClient.Template, db.Querying)
		if err != nil {
			return nil, err
		}

		return mod, nil
	}

	return nil, nil
}
