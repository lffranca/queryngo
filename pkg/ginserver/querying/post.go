package querying

import (
	"github.com/gin-gonic/gin"
	module "github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/postgres"
	"log"
	"net/http"
)

func queryingPOST(db *postgres.Client) gin.HandlerFunc {
	format, err := formatter.New()
	if err != nil {
		log.Panicln(err)
	}

	mod, err := module.New(db.Template, format.Template, db.Querying)
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

		var body interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Println("[WARNING] c.ShouldBindJSON: ", err)
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
