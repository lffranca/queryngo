package gaws

import (
	"errors"
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
