package importdata

import "context"

type ProcessFileRepository interface {
	ProducerProcessFile(ctx context.Context, value []byte) error
}
