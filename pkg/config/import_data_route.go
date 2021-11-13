package config

type ImportDataRoute struct {
	Enabled  bool   `json:"enabled" yaml:"enabled"`
	Bucket   string `json:"bucket" yaml:"bucket"`
	BasePath string `json:"base_path" yaml:"base_path"`
}
