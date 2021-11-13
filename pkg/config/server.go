package config

type Server struct {
	Prefix   string `json:"prefix" yaml:"prefix"`
	Port     string `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
	Routes   Routes `json:"routes" yaml:"routes"`
}
