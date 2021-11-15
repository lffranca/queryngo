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
	Prefix               *string
	Routes               *RoutesOptions
	QueryingRepository   QueryingRepository
	ImportDataRepository ImportDataRepository
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

		if pkg.ImportDataRepository == nil {
			return errors.New("ImportDataRepository is required in import data route")
		}
	}

	if pkg.Routes.Querying.Enabled {
		if pkg.QueryingRepository == nil {
			return errors.New("QueryingRepository is required in querying route")
		}
	}

	return nil
}
