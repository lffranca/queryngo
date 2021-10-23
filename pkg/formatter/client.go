package formatter

func New() (*Client, error) {
	client := new(Client)
	client.common.client = client
	client.Template = (*TemplateService)(&client.common)

	return client, nil
}

type service struct {
	client *Client
}

type Client struct {
	common   service
	Template *TemplateService
}
