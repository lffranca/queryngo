package server

import "errors"

type Options struct {
	QueryingRepository  QueryingRepository
	FormatterRepository FormatterRepository
	TemplateRepository  TemplateRepository
}

func (pkg *Options) validate() error {
	if pkg.QueryingRepository == nil {
		return errors.New("QueryingRepository is required")
	}

	if pkg.FormatterRepository == nil {
		return errors.New("FormatterRepository is required")
	}

	if pkg.TemplateRepository == nil {
		return errors.New("TemplateRepository is required")
	}

	return nil
}
