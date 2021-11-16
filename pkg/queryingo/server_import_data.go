package queryingo

import (
	"github.com/lffranca/queryngo/domain/importdata"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/gkafka"
	"github.com/lffranca/queryngo/pkg/guuid"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func serverImportDataRoute(client config.Server, db *postgres.Client, broker *gkafka.Server) (*importdata.ImportData, error) {
	if client.Routes.ImportData.Enabled {
		awsClient, err := gaws.New(&client.Routes.ImportData.Bucket)
		if err != nil {
			return nil, err
		}

		uuidClient, err := guuid.New()
		if err != nil {
			return nil, err
		}

		//mod, err := importdata.New(awsClient.Storage, db.File, uuidClient.GenerateImportData)
		mod, err := importdata.New(&importdata.Options{
			Prefix:                &client.Prefix,
			Bucket:                &client.Routes.ImportData.Bucket,
			BasePath:              &client.Routes.ImportData.BasePath,
			StorageRepository:     awsClient.Storage,
			FileRepository:        db.File,
			UUIDRepository:        uuidClient.GenerateImportData,
			ProcessFileRepository: broker.ProcessFile,
		})
		if err != nil {
			return nil, err
		}

		return mod, nil
	}

	return nil, nil
}
