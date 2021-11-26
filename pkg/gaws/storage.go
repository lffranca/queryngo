package gaws

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lffranca/queryngo/domain"
	"github.com/lffranca/queryngo/pkg/gaws/model"
	"io"
)

type StorageService service

func (pkg *StorageService) Download(ctx context.Context, key *string) (io.Reader, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	buffW, err := model.NewBufferWriterAt(buff)
	if err != nil {
		return nil, err
	}

	downloader := manager.NewDownloader(client)
	if _, err := downloader.Download(ctx, buffW, &s3.GetObjectInput{
		Bucket: pkg.client.bucket,
		Key:    key,
	}); err != nil {
		return nil, err
	}

	return buff, nil
}

func (pkg *StorageService) PreSign(ctx context.Context, key, contentType *string) (*string, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	psClient := s3.NewPresignClient(client)
	res, err := psClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket:              pkg.client.bucket,
		Key:                 key,
		ResponseContentType: contentType,
	})
	if err != nil {
		return nil, err
	}

	return &res.URL, nil
}

func (pkg *StorageService) ListObjects(ctx context.Context) ([]*domain.FileInfo, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: pkg.client.bucket,
	})
	if err != nil {
		return nil, err
	}

	var files []*domain.FileInfo
	for _, object := range output.Contents {
		size := int(object.Size)
		files = append(files, &domain.FileInfo{
			Key:          object.Key,
			Size:         &size,
			LastModified: object.LastModified,
		})
	}

	return files, nil
}

func (pkg *StorageService) Upload(ctx context.Context, key, contentType *string, data io.Reader) error {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return err
	}

	uploader := manager.NewUploader(client)
	if _, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Key:         key,
		ContentType: contentType,
		Body:        data,
		Bucket:      pkg.client.bucket,
	}); err != nil {
		return err
	}

	return nil
}

func (pkg *StorageService) clientS3(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
