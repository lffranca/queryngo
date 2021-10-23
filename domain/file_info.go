package domain

import (
	"time"
)

type FileInfo struct {
	Name         *string
	Extension    *string
	Key          *string
	Path         *string
	Size         *int
	ContentType  *string
	LastModified *time.Time
}
