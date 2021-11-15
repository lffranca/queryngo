package domain

type FileInfoResult struct {
	ID          *int      `json:"id,omitempty" yaml:"id,omitempty"`
	Sheet       *string   `json:"sheet,omitempty" yaml:"sheet,omitempty"`
	Path        *string   `json:"path,omitempty" yaml:"path,omitempty"`
	Extension   *string   `json:"extension,omitempty" yaml:"extension,omitempty"`
	ContentType *string   `json:"content_type,omitempty" yaml:"content_type,omitempty"`
	Columns     []*string `json:"columns,omitempty" yaml:"columns,omitempty"`
	Key         *string   `json:"key,omitempty" yaml:"key,omitempty"`
	Size        *int      `json:"size,omitempty" yaml:"size,omitempty"`
	ParentID    *int      `json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
}
