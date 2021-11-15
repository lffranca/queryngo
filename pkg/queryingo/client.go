package queryingo

import (
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/postgres"
	"log"
	"sync"
)

func clientRun(wgParent *sync.WaitGroup, client config.Server) {
	defer wgParent.Done()

	// database connection in principal client thread
	db, err := postgres.New(&client.Database)
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	// kafka connection in principal client thread
	broker, err := newBroker(db, client)
	if err != nil {
		log.Panicln(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go serverClientRun(wg, client, db)

	if broker != nil {
		wg.Add(1)
		go brokerProcessedFileListener(wg, broker)
	}

	wg.Wait()
}
