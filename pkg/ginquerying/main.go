package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	module "github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/formatter"
	"github.com/lffranca/queryngo/pkg/postgres"
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

	format, err := formatter.New()
	if err != nil {
		log.Panicln(err)
	}

	mod, err := module.New(db.Template, format.Template, db.Querying)
	if err != nil {
		log.Panicln(err)
	}

	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
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
	})

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		log.Panicln(err)
	}
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
