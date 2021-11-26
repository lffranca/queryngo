package queryingo

import (
	"github.com/lffranca/queryngo/domain/file"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func serverFile(client config.Server, db *postgres.Client, awsClient *gaws.Client, formatterClient *formatter.Client) (*file.File, error) {
	if client.Routes.ImportData.Enabled {
		mod, err := file.New(&file.Options{
			Repository:          db.File,
			ProcessedRepository: db.FileProcessed,
			ConfigRepository:    db.FileConfig,
			StorageRepository:   awsClient.Storage,
			ReaderRepository:    formatterClient.CSV,
		})
		if err != nil {
			return nil, err
		}

		return mod, nil
	}

	return nil, nil
}
