package queryingo

import (
	"github.com/lffranca/queryngo/domain/processedfile"
	"github.com/lffranca/queryngo/pkg/config"
	"github.com/lffranca/queryngo/pkg/gkafka"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func newBroker(db *postgres.Client, client config.Server) (*gkafka.Server, error) {
	if client.Broker.Brokers == nil {
		return nil, nil
	}

	mod, err := processedfile.New(db.FileProcessed, db.File)
	if err != nil {
		return nil, err
	}

	return gkafka.New(&gkafka.Options{
		Brokers:                 client.Broker.Brokers,
		Network:                 "tcp",
		ProcessFileTopic:        "process-file",
		ProcessedFileTopic:      client.Broker.BrokerListeners.ProcessedFileTopic,
		ProcessedFileRepository: mod,
	})
}
