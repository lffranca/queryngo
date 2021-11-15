package queryingo

import (
	"context"
	"github.com/lffranca/queryngo/pkg/gkafka"
	"log"
	"sync"
)

func brokerProcessedFileListener(wgParent *sync.WaitGroup, broker *gkafka.Server) {
	defer wgParent.Done()

	ctx := context.Background()
	if err := broker.ProcessFile.ConsumerProcessedFile(ctx); err != nil {
		log.Panicln(err)
	}
}
