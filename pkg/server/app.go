package server

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func New(options *Options) (*Server, error) {
	if options == nil {
		return nil, errors.New("options is required")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	server := new(Server)
	server.common.Server = server
	server.Querying = (*QueryingService)(&server.common)
	server.ImportData = (*ImportDataService)(&server.common)

	server.app = gin.Default()
	server.routes()

	return server, nil
}

type Server struct {
	common              service
	app                 *gin.Engine
	Querying            *QueryingService
	ImportData          *ImportDataService
	QueryingRepository  QueryingRepository
	FormatterRepository FormatterRepository
	TemplateRepository  TemplateRepository
}

func (pkg *Server) routes() {
	v1 := pkg.app.Group("/v1")
	{
		querying := v1.Group("/querying")
		{
			querying.POST("", pkg.Querying.queryingPOST())
		}
	}
}

func (pkg *Server) Run(addr ...string) error {
	return pkg.app.Run(addr...)
}

type service struct {
	Server *Server
}
