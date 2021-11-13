package main

import (
	"fmt"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/guuid"
	"github.com/lffranca/queryngo/pkg/postgres"
	"github.com/lffranca/queryngo/pkg/server"
	"log"
	"sync"
)

func main() {
	conf, err := config.New(nil)
	if err != nil {
		log.Panicln(err)
	}

	wg := &sync.WaitGroup{}
	for _, client := range conf.Servers {
		wg.Add(1)
		go newClientServer(wg, client)
	}

	wg.Wait()
}

func newClientServer(wgParent *sync.WaitGroup, client config.Server) {
	defer wgParent.Done()

	db, err := postgres.New(&client.Database)
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	var formatterRepository *formatter.TemplateService
	if client.Routes.Querying.Enabled {
		formatterClient, err := formatter.New()
		if err != nil {
			log.Panicln(err)
		}

		formatterRepository = formatterClient.Template
	}

	var storageRepository *gaws.StorageService
	var uuidRepository *guuid.GenerateService
	if client.Routes.ImportData.Enabled {
		awsClient, err := gaws.New(&client.Routes.ImportData.Bucket)
		if err != nil {
			log.Panicln(err)
		}

		storageRepository = awsClient.Storage

		uuidClient, err := guuid.New()
		if err != nil {
			log.Panicln(err)
		}

		uuidRepository = uuidClient.GenerateImportData
	}

	app, err := server.New(&server.Options{
		Prefix:              &client.Prefix,
		QueryingRepository:  db.Querying,
		TemplateRepository:  db.Template,
		FormatterRepository: formatterRepository,
		StorageRepository:   storageRepository,
		FileRepository:      db.File,
		UUIDRepository:      uuidRepository,
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
