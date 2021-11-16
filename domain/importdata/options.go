package importdata

import "errors"

type Options struct {
	Prefix                *string
	Bucket                *string
	BasePath              *string
	StorageRepository     StorageRepository
	FileRepository        FileRepository
	UUIDRepository        UUIDRepository
	ProcessFileRepository ProcessFileRepository
}

func (options *Options) validate() error {
	if options.Prefix == nil {
		return errors.New("prefix param is required")
	}

	if options.Bucket == nil {
		return errors.New("bucket param is required")
	}

	if options.BasePath == nil {
		return errors.New("base path param is required")
	}

	if options.StorageRepository == nil {
		return errors.New("StorageRepository param is required")
	}

	if options.FileRepository == nil {
		return errors.New("FileRepository param is required")
	}

	if options.UUIDRepository == nil {
		return errors.New("UUIDRepository param is required")
	}

	if options.ProcessFileRepository == nil {
		return errors.New("ProcessFileRepository param is required")
	}

	return nil
}
