package guuid

import (
	"github.com/google/uuid"
)

func New() (*Client, error) {
	client := new(Client)
	client.common.client = client
	client.GenerateImportData = (*GenerateService)(&client.common)

	return client, nil
}

type service struct {
	client *Client
}

type Client struct {
	common             service
	GenerateImportData *GenerateService
}

func (pkg *Client) uuid() (*string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	ids := id.String()

	return &ids, nil
}
