package file

import "errors"

type Options struct {
	Repository          Repository
	ProcessedRepository ProcessedRepository
	ConfigRepository    ConfigRepository
	StorageRepository   StorageRepository
	ReaderRepository    ReaderRepository
}

func (pkg *Options) validate() error {
	if pkg.Repository == nil {
		return errors.New("repository param is required")
	}

	if pkg.ProcessedRepository == nil {
		return errors.New("ProcessedRepository param is required")
	}

	if pkg.ConfigRepository == nil {
		return errors.New("ConfigRepository param is required")
	}

	if pkg.StorageRepository == nil {
		return errors.New("StorageRepository param is required")
	}

	if pkg.ReaderRepository == nil {
		return errors.New("ReaderRepository param is required")
	}

	return nil
}
