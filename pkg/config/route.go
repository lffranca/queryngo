package config

type Routes struct {
	ImportData ImportDataRoute `json:"import_data" yaml:"import_data"`
	Querying   QueryingRoute   `json:"querying" yaml:"querying"`
}
