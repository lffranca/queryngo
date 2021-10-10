package querying

import "context"

type Formatter interface {
	Transform(ctx context.Context, template []byte, input interface{}) (expected []byte, err error)
}
