package querying

import (
	"github.com/gin-gonic/gin"
	module "github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/postgres"
	"log"
	"net/http"
)

func queryingMultiTenantPOST(db *postgres.Client) gin.HandlerFunc {
	format, err := formatter.New()
	if err != nil {
		log.Panicln(err)
	}

	tenant, err := postgres.NewMultiTenant(db.DB())
	if err != nil {
		log.Panicln(err)
	}

	return func(c *gin.Context) {
		var query queryStringBind
		if err := c.ShouldBindQuery(&query); err != nil {
			log.Println("c.ShouldBindQuery: ", err)
			c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}

		var headers headerBind
		if err := c.ShouldBindHeader(&headers); err != nil {
			log.Println("c.ShouldBindHeader: ", err)
			c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}

		var body interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Println("[WARNING] c.ShouldBindJSON: ", err)
		}

		tenantDB, err := tenant.Client(c.Request.Context(), &headers.Sub)
		if err != nil {
			log.Println("tenant.Client: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		mod, err := module.New(db.Template, format.Template, tenantDB.Querying)
		if err != nil {
			log.Println("module.New: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		data, err := mod.Execute(c.Request.Context(), query.QueryID, query.FormatID, body)
		if err != nil {
			log.Println("mod.Execute: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/json", data)
	}
}
