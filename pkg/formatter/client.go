package formatter

func New() (*Client, error) {
	client := new(Client)
	client.common.client = client
	client.Template = (*TemplateService)(&client.common)
	client.CSV = (*CSVService)(&client.common)

	return client, nil
}

type Client struct {
	common   service
	Template *TemplateService
	CSV      *CSVService
}

type service struct {
	client *Client
}
