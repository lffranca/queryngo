package config

type Broker struct {
	Brokers         []string        `json:"brokers,omitempty" yaml:"brokers,omitempty"`
	BrokerListeners BrokerListeners `json:"listeners,omitempty" yaml:"listeners,omitempty"`
}
