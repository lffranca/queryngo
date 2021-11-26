package model

import (
	"errors"
	"io"
)

func NewBufferWriterAt(w io.Writer) (*BufferWriterAt, error) {
	if w == nil {
		return nil, errors.New("w param is required")
	}

	buff := new(BufferWriterAt)
	buff.w = w

	return buff, nil
}

type BufferWriterAt struct {
	w io.Writer
}

func (fw BufferWriterAt) WriteAt(p []byte, offset int64) (n int, err error) {
	return fw.w.Write(p)
}
