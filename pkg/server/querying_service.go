package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/server/presenter"
	"log"
	"net/http"
)

type QueryingService service

func (pkg *QueryingService) queryingPOST() gin.HandlerFunc {
	mod, err := querying.New(
		pkg.Server.TemplateRepository,
		pkg.Server.FormatterRepository,
		pkg.Server.QueryingRepository)
	if err != nil {
		log.Panicln(err)
	}

	return func(c *gin.Context) {
		var query presenter.QueryingURI
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
