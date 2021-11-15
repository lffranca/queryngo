package processedfile

import (
	"context"
	"github.com/lffranca/queryngo/domain"
)

type FileRepository interface {
	Update(ctx context.Context, item *domain.FileInfo) error
}
