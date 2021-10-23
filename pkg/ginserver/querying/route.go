package querying

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/pkg/postgres"
)

func Route(app *gin.RouterGroup, db *postgres.Client) {
	app.POST("", queryingPOST(db))
}

func RouteMultiTenant(app *gin.RouterGroup, db *postgres.Client) {
	app.POST("", queryingMultiTenantPOST(db))
}
