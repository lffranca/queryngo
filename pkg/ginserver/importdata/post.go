package importdata

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/domain/importdata"
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/guuid"
	"github.com/lffranca/queryngo/pkg/postgres"
	"log"
	"net/http"
)

func importDataPOST(db *postgres.Client, awsClient *gaws.Client) gin.HandlerFunc {
	generate, err := guuid.New()
	if err != nil {
		log.Panicln(err)
	}

	tenant, err := postgres.NewMultiTenant(db.DB())
	if err != nil {
		log.Panicln(err)
	}

	return func(c *gin.Context) {
		var headers headerBind
		if err := c.ShouldBindHeader(&headers); err != nil {
			log.Println("c.ShouldBindHeader: ", err)
			c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}

		var fileItem formBody
		if err := c.ShouldBind(&fileItem); err != nil {
			log.Println("c.ShouldBind: ", err)
			c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}

		tenantDB, err := tenant.Client(c.Request.Context(), &headers.Sub)
		if err != nil {
			log.Println("tenant.Client: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		basePath := "upload/files"
		mod, err := importdata.New(awsClient.StorageImportData, tenantDB.SaveFileKeyImportData, generate.GenerateImportData, &basePath)
		if err != nil {
			log.Println("importData.New: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		file, err := fileItem.File.Open()
		if err != nil {
			log.Println("fileItem.File.Open: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		fileName := fileItem.File.Filename

		var contentType *string
		if value, ok := fileItem.File.Header["Content-Type"]; ok {
			if len(value) > 0 {
				contentType = &value[0]
			}
		}

		fileSize := int(fileItem.File.Size)
		if err := mod.Import(c.Request.Context(), &fileName, contentType, &fileSize, file); err != nil {
			log.Println("mod.Import: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
