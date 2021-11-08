package server

import "context"

type FormatterRepository interface {
	Transform(ctx context.Context, template []byte, input interface{}) (expected []byte, err error)
}
