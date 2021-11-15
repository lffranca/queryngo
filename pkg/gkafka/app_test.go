package gkafka

import (
	"context"
	"github.com/lffranca/queryngo/domain"
	"testing"
)

func TestProcessFileService_ConsumerProcessedFile(t *testing.T) {
	app, err := New(&Options{
		Brokers:            []string{"localhost:9092"},
		ProcessFileTopic:   "process-file",
		ProcessedFileTopic: "blinclass-processed-file",
		Network:            "tcp",
	})
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()

	if err := app.ProcessFile.ConsumerProcessedFile(ctx); err != nil {
		t.Error(err)
		return
	}
}

type mockProcessedFile struct{}

func (pkg *mockProcessedFile) ProcessedFileResult(ctx context.Context, info *domain.FileInfo) error {
	return nil
}

func TestProcessFileService_ProducerProcessFile(t *testing.T) {
	app, err := New(&Options{
		Brokers:                 []string{"localhost:9092"},
		ProcessFileTopic:        "process-file",
		ProcessedFileTopic:      "blinclass-processed-file",
		Network:                 "tcp",
		ProcessedFileRepository: &mockProcessedFile{},
	})
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()

	itemString := `{"id": 8, "key": "81749ba4-44b1-11ec-a5bc-641c67ac6c21.xlsx", "path": "81749ba4-44b1-11ec-a5bc-641c67ac6c21.xlsx", "name": "Historico de vendas - Template.xlsx", "extension": ".xlsx", "size": 5292, "content_type": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "last_modified": "2021-11-15T22:03:13.130Z", "prefix": "blinclass", "bucket": "upowl-client-blinclass"}`

	if err := app.ProcessFile.ProducerProcessFile(ctx, []byte(itemString)); err != nil {
		t.Error(err)
		return
	}
}
