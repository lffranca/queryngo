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
	server.Prefix = options.Prefix
	server.RoutesOptions = options.Routes
	server.Querying = (*QueryingService)(&server.common)
	server.ImportData = (*ImportDataService)(&server.common)

	server.QueryingRepository = options.QueryingRepository
	server.ImportDataRepository = options.ImportDataRepository

	server.app = gin.Default()
	server.routes()

	return server, nil
}

type Server struct {
	common               service
	app                  *gin.Engine
	Prefix               *string
	RoutesOptions        *RoutesOptions
	Querying             *QueryingService
	ImportData           *ImportDataService
	ImportDataRepository ImportDataRepository
	QueryingRepository   QueryingRepository
}

func (pkg *Server) routes() {
	v1 := pkg.app.Group("/v1")
	{
		if pkg.RoutesOptions.Querying.Enabled {
			querying := v1.Group("/querying")
			{
				querying.POST("", pkg.Querying.queryingPOST)
			}
		}

		if pkg.RoutesOptions.ImportData.Enabled {
			importData := v1.Group("/import-data")
			{
				importData.POST("", pkg.ImportData.importDataPOST)
			}
		}
	}
}

func (pkg *Server) Run(addr ...string) error {
	return pkg.app.Run(addr...)
}

type service struct {
	Server *Server
}
