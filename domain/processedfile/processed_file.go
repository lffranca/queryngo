package processedfile

import (
	"context"
	"errors"
	"github.com/lffranca/queryngo/domain"
)

func New(fileProcessed FileProcessedRepository, file FileRepository) (*processedFile, error) {
	if fileProcessed == nil {
		return nil, errors.New("fileProcessed is required param")
	}

	if file == nil {
		return nil, errors.New("file is required param")
	}

	client := new(processedFile)
	client.fileProcessed = fileProcessed
	client.file = file
	return client, nil
}

type processedFile struct {
	fileProcessed FileProcessedRepository
	file          FileRepository
}

func (pkg *processedFile) ProcessedFileResult(ctx context.Context, info *domain.FileInfo) error {
	for index, item := range info.Results {
		item.ParentID = info.ID
		item.Bucket = info.Bucket
		item.Prefix = info.Prefix
		info.Results[index] = item
	}

	if err := pkg.fileProcessed.SaveAll(ctx, info.Results); err != nil {
		return err
	}

	info.Status = domain.FileStatusProcessed

	if err := pkg.file.Update(ctx, info); err != nil {
		return err
	}

	return nil
}
