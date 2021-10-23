package importdata

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func RouteMultiTenant(app *gin.RouterGroup, db *postgres.Client, awsClient *gaws.Client) {
	app.POST("", importDataPOST(db, awsClient))
}
