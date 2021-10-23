package ginserver

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/ginserver/importdata"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func New(db *postgres.Client, awsClient *gaws.Client, port *string) (*Server, error) {
	if db == nil || awsClient == nil || port == nil {
		return nil, errors.New("invalid params")
	}

	server := new(Server)
	server.db = db
	server.aws = awsClient
	server.port = port
	server.app = gin.Default()

	server.routes()

	return server, nil
}

type Server struct {
	db   *postgres.Client
	aws  *gaws.Client
	port *string
	app  *gin.Engine
}

func (pkg *Server) routes() {
	v1 := pkg.app.Group("/v1")
	{
		multiTenant := v1.Group("/multi-tenancy")
		{
			importdata.Route(multiTenant.Group("/import-data"), pkg.db, pkg.aws)
		}
	}
}

func (pkg *Server) Run() error {
	return pkg.app.Run(fmt.Sprintf(":%s", *pkg.port))
}
