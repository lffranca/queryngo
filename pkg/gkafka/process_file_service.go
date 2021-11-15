package gkafka

import (
	"context"
	"encoding/json"
	"github.com/lffranca/queryngo/domain"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

type ProcessFileService service

func (pkg *ProcessFileService) ConsumerProcessedFile(ctx context.Context) error {
	conn, err := kafka.DialLeader(ctx,
		pkg.Server.network,
		strings.Join(pkg.Server.brokers, ","),
		pkg.Server.processedFileTopic, 0)
	if err != nil {
		return err
	}

	_, offsetLast, err := conn.ReadOffsets()
	if err != nil {
		return err
	}

	if err := conn.Close(); err != nil {
		return err
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   pkg.Server.brokers,
		Topic:     pkg.Server.processedFileTopic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	if err := r.SetOffset(offsetLast); err != nil {
		return err
	}

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		var item *domain.FileInfo
		if err := json.Unmarshal(m.Value, &item); err != nil {
			log.Println("json.Unmarshal: ", err)
			continue
		}

		if err := pkg.Server.processedFileRepository.ProcessedFileResult(ctx, item); err != nil {
			log.Println("ProcessedFileResult: ", err)
			continue
		}
	}

	if err := r.Close(); err != nil {
		return err
	}

	return nil
}

func (pkg *ProcessFileService) ProducerProcessFile(ctx context.Context, value []byte) error {
	conn, err := kafka.DialLeader(ctx,
		pkg.Server.network,
		strings.Join(pkg.Server.brokers, ","),
		pkg.Server.processFileTopic, 0)
	if err != nil {
		return err
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err = conn.Write(value); err != nil {
		return err
	}

	return nil
}
