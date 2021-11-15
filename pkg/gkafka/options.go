package gkafka

import "errors"

type Options struct {
	Brokers                 []string
	Network                 string
	ProcessFileTopic        string
	ProcessedFileTopic      string
	ProcessedFileRepository ProcessedFileRepository
}

func (pkg *Options) validate() error {
	if len(pkg.Brokers) <= 0 {
		return errors.New("brokers param is required")
	}

	if len(pkg.Network) <= 0 {
		return errors.New("network param is required")
	}

	if len(pkg.ProcessFileTopic) <= 0 {
		return errors.New("process file topic param is required")
	}

	if len(pkg.ProcessedFileTopic) <= 0 {
		return errors.New("process file topic param is required")
	}

	if pkg.ProcessedFileRepository == nil {
		return errors.New("ProcessedFileRepository param is required")
	}

	return nil
}
