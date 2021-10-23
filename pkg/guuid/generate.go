package guuid

import "context"

type GenerateService service

func (pkg *GenerateService) UUID(ctx context.Context) (*string, error) {
	return pkg.client.uuid()
}
