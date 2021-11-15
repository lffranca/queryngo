package queryingo

import (
	"github.com/lffranca/queryngo/domain/importdata"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/guuid"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func serverImportDataRoute(route config.ImportDataRoute, db *postgres.Client) (*importdata.ImportData, error) {
	if route.Enabled {
		awsClient, err := gaws.New(&route.Bucket)
		if err != nil {
			return nil, err
		}

		uuidClient, err := guuid.New()
		if err != nil {
			return nil, err
		}

		mod, err := importdata.New(awsClient.Storage, db.File, uuidClient.GenerateImportData)
		if err != nil {
			return nil, err
		}

		return mod, nil
	}

	return nil, nil
}
