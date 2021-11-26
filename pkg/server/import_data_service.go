package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/pkg/server/presenter"
	"log"
	"net/http"
)

type ImportDataService service

func (pkg *ImportDataService) importDataPOST(c *gin.Context) {
	var fileItem presenter.ImportDataBody
	if err := c.ShouldBind(&fileItem); err != nil {
		log.Println("c.ShouldBind: ", err)
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
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
	if err := pkg.Server.importDataRepository.Import(c.Request.Context(), &fileName, contentType, &fileSize, file); err != nil {
		log.Println("mod.Import: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
