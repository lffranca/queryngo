package domain

import (
	"time"
)

type FileInfo struct {
	ID           *int              `json:"id,omitempty" yaml:"id,omitempty"`
	Name         *string           `json:"name,omitempty" yaml:"name,omitempty"`
	Extension    *string           `json:"extension,omitempty" yaml:"extension,omitempty"`
	Key          *string           `json:"key,omitempty" yaml:"key,omitempty"`
	Path         *string           `json:"path,omitempty" yaml:"path,omitempty"`
	Size         *int              `json:"size,omitempty" yaml:"size,omitempty"`
	ContentType  *string           `json:"content_type,omitempty" yaml:"content_type,omitempty"`
	LastModified *time.Time        `json:"last_modified,omitempty" yaml:"last_modified,omitempty"`
	Prefix       *string           `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Bucket       *string           `json:"bucket,omitempty" yaml:"bucket,omitempty"`
	Results      []*FileInfoResult `json:"result,omitempty" yaml:"result,omitempty"`
	Status       FileStatus        `json:"status,omitempty" yaml:"status,omitempty"`
}
