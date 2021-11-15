package config

type BrokerListeners struct {
	ProcessedFileTopic string `json:"processed_file_topic,omitempty" yaml:"processed_file_topic,omitempty"`
}
