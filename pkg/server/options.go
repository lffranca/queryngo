package server

import "errors"

type ImportDataRouteOptions struct {
	Enabled  bool
	Bucket   string
	BasePath string
}

type QueryingRouteOptions struct {
	Enabled bool
}

type RoutesOptions struct {
	ImportData ImportDataRouteOptions
	Querying   QueryingRouteOptions
}

type Options struct {
	Prefix              *string
	Routes              *RoutesOptions
	QueryingRepository  QueryingRepository
	FormatterRepository FormatterRepository
	TemplateRepository  TemplateRepository
	StorageRepository   StorageRepository
	FileRepository      FileRepository
	UUIDRepository      UUIDRepository
}

func (pkg *Options) validate() error {
	if pkg.Routes == nil {
		return errors.New("routes param is required")
	}

	if pkg.Routes.ImportData.Enabled {
		if len(pkg.Routes.ImportData.Bucket) <= 0 {
			return errors.New("bucket param is required in import data route")
		}

		if len(pkg.Routes.ImportData.BasePath) <= 0 {
			return errors.New("base path is required in import data route")
		}

		if pkg.StorageRepository == nil {
			return errors.New("StorageRepository is required in import data route")
		}

		if pkg.FileRepository == nil {
			return errors.New("FileRepository is required in import data route")
		}

		if pkg.UUIDRepository == nil {
			return errors.New("UUIDRepository is required in import data route")
		}
	}

	if pkg.Routes.Querying.Enabled {
		if pkg.QueryingRepository == nil {
			return errors.New("QueryingRepository is required in querying route")
		}

		if pkg.FormatterRepository == nil {
			return errors.New("FormatterRepository is required in querying route")
		}

		if pkg.TemplateRepository == nil {
			return errors.New("TemplateRepository is required in querying route")
		}
	}

	return nil
}
