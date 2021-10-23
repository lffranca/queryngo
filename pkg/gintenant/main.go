package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	module "github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/postgres"
	"github.com/lffranca/queryngo/repository/format"
	"github.com/lffranca/queryngo/repository/formatter"
	"log"
	"net/http"
	"os"
)

func main() {
	connString := os.Getenv("DB_CONN_STRING")
	db, err := postgres.New(&connString)
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	tenant, err := postgres.NewMultiTenant(db.DB())
	if err != nil {
		log.Panicln(err)
	}

	formatRepository, err := format.New(db)
	if err != nil {
		log.Panicln(err)
	}

	formatterRepository := formatter.NewTemplate()

	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
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

		mod, err := module.New(formatRepository, formatterRepository, tenantDB.Querying)
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
	})

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		log.Panicln(err)
	}
}

type headerBind struct {
	Sub string `header:"X-Sub" binding:"required"`
}

type queryStringBind struct {
	QueryID  *string `form:"query_id" json:"query_id" binding:"required"`
	FormatID *string `form:"format_id" json:"format_id" binding:"required"`
}

func init() {
	envs := []string{
		"API_PORT",
		"DB_CONN_STRING",
	}

	for _, env := range envs {
		if _, ok := os.LookupEnv(env); !ok {
			log.Panicf("env var is required: %s\n", env)
		}
	}
}
