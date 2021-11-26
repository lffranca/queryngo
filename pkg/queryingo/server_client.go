package queryingo

import (
	"fmt"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/gkafka"
	"github.com/lffranca/queryngo/pkg/postgres"
	"github.com/lffranca/queryngo/pkg/server"
	"log"
	"sync"
)

func serverClientRun(wgParent *sync.WaitGroup, client config.Server, db *postgres.Client, broker *gkafka.Server) {
	defer wgParent.Done()

	awsClient, err := gaws.New(&client.Routes.ImportData.Bucket)
	if err != nil {
		log.Panicln(err)
	}

	formatterClient, err := formatter.New()
	if err != nil {
		log.Panicln(err)
	}

	queryingMod, err := serverQueryingRoute(client.Routes.Querying, db, formatterClient)
	if err != nil {
		log.Panicln(err)
	}

	importDataMod, err := serverImportDataRoute(client, db, broker, awsClient)
	if err != nil {
		log.Panicln(err)
	}

	storageMod, err := serverFile(client, db, awsClient, formatterClient)
	if err != nil {
		log.Panicln(err)
	}

	app, err := server.New(&server.Options{
		Prefix:               &client.Prefix,
		QueryingRepository:   queryingMod,
		ImportDataRepository: importDataMod,
		StorageRepository:    storageMod,
		Routes: &server.RoutesOptions{
			Querying: server.QueryingRouteOptions{
				Enabled: client.Routes.Querying.Enabled,
			},
			ImportData: server.ImportDataRouteOptions{
				Enabled:  client.Routes.ImportData.Enabled,
				Bucket:   client.Routes.ImportData.Bucket,
				BasePath: client.Routes.ImportData.BasePath,
			},
		},
	})

	if err := app.Run(fmt.Sprintf(":%s", client.Port)); err != nil {
		log.Panicln(err)
	}
}
