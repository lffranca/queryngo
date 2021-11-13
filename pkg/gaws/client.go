package gaws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lffranca/queryngo/domain"
	"io"
)

func New(bucket *string) (*Client, error) {
	if bucket == nil {
		return nil, errors.New("invalid params")
	}

	client := new(Client)
	client.bucket = bucket
	client.common.client = client
	client.Storage = (*StorageService)(&client.common)

	return client, nil
}

type service struct {
	client *Client
}

type Client struct {
	bucket  *string
	common  service
	Storage *StorageService
}

func (pkg *Client) upload(ctx context.Context, key, contentType *string, data io.Reader) error {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return err
	}

	uploader := manager.NewUploader(client)
	if _, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Key:         key,
		ContentType: contentType,
		Body:        data,
		Bucket:      pkg.bucket,
	}); err != nil {
		return err
	}

	return nil
}

func (pkg *Client) listObjects(ctx context.Context) ([]*domain.FileInfo, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: pkg.bucket,
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

func (pkg *Client) preSign(ctx context.Context, key, contentType *string) (*string, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	psClient := s3.NewPresignClient(client)
	res, err := psClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket:              pkg.bucket,
		Key:                 key,
		ResponseContentType: contentType,
	})
	if err != nil {
		return nil, err
	}

	return &res.URL, nil
}

func (pkg *Client) clientS3(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
