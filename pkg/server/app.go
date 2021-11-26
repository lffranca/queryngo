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
	server.prefix = options.Prefix
	server.routesOptions = options.Routes
	server.querying = (*QueryingService)(&server.common)
	server.importData = (*ImportDataService)(&server.common)
	server.storage = (*StorageService)(&server.common)

	server.queryingRepository = options.QueryingRepository
	server.importDataRepository = options.ImportDataRepository
	server.storageRepository = options.StorageRepository

	server.app = gin.Default()
	server.routes()

	return server, nil
}

type Server struct {
	common               service
	app                  *gin.Engine
	prefix               *string
	routesOptions        *RoutesOptions
	querying             *QueryingService
	importData           *ImportDataService
	storage              *StorageService
	importDataRepository ImportDataRepository
	queryingRepository   QueryingRepository
	storageRepository    StorageRepository
}

func (pkg *Server) routes() {
	v1 := pkg.app.Group("/v1")
	{
		if pkg.routesOptions.Querying.Enabled {
			querying := v1.Group("/querying")
			{
				querying.POST("", pkg.querying.queryingPOST)
			}
		}

		if pkg.routesOptions.ImportData.Enabled {
			importData := v1.Group("/import-data")
			{
				importData.POST("", pkg.importData.importDataPOST)
			}

			storage := v1.Group("/storage")
			{
				storage.GET("/config", pkg.storage.listConfigGET)
				storage.DELETE("/config", pkg.storage.configDELETE)
				storage.POST("/config", pkg.storage.configPOST)
				storage.GET("/content", pkg.storage.fileContentGET)
				storage.GET("/processed", pkg.storage.listProcessedGET)
				storage.GET("", pkg.storage.listGET)
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
