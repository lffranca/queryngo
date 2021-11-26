package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/queryngo/pkg/server/presenter"
	"log"
	"net/http"
)

type StorageService service

func (pkg *StorageService) listConfigGET(c *gin.Context) {
	var queryParent presenter.ParentURI
	if err := c.ShouldBindQuery(&queryParent); err != nil {
		log.Println("c.ShouldBindQuery queryParent: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var queryList presenter.ListCommonURI
	if err := c.ShouldBindQuery(&queryList); err != nil {
		log.Println("c.ShouldBindQuery queryList: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	items, err := pkg.Server.storageRepository.ListConfig(
		c.Request.Context(),
		queryParent.ID,
		queryList.Offset,
		queryList.Limit,
		queryList.Search,
	)
	if err != nil {
		log.Println("storageRepository.ListConfig: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (pkg *StorageService) configDELETE(c *gin.Context) {
	var queryParent presenter.ParentURI
	if err := c.ShouldBindQuery(&queryParent); err != nil {
		log.Println("c.ShouldBindQuery queryParent: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := pkg.Server.storageRepository.DeleteConfig(c.Request.Context(), queryParent.ID); err != nil {
		log.Println("storageRepository.DeleteConfig: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (pkg *StorageService) configPOST(c *gin.Context) {
	var body presenter.FileConfigBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("c.ShouldBindJSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	item, err := pkg.Server.storageRepository.SaveConfig(c.Request.Context(), body.Entity())
	if err != nil {
		log.Println("storageRepository.SaveConfig: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (pkg *StorageService) fileContentGET(c *gin.Context) {
	var queryParent presenter.ParentURI
	if err := c.ShouldBindQuery(&queryParent); err != nil {
		log.Println("c.ShouldBindQuery queryParent: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	content, err := pkg.Server.storageRepository.ProcessedFileContent(
		c.Request.Context(),
		queryParent.ID,
	)
	if err != nil {
		log.Println("storageRepository.ProcessedFileContent: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (pkg *StorageService) listProcessedGET(c *gin.Context) {
	var queryParent presenter.ParentURI
	if err := c.ShouldBindQuery(&queryParent); err != nil {
		log.Println("c.ShouldBindQuery queryParent: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var queryList presenter.ListCommonURI
	if err := c.ShouldBindQuery(&queryList); err != nil {
		log.Println("c.ShouldBindQuery queryList: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	items, err := pkg.Server.storageRepository.ListProcessedFile(
		c.Request.Context(),
		queryParent.ID,
		queryList.Offset,
		queryList.Limit,
		queryList.Search,
	)
	if err != nil {
		log.Println("storageRepository.ListProcessedFile: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (pkg *StorageService) listGET(c *gin.Context) {
	var queryList presenter.ListCommonURI
	if err := c.ShouldBindQuery(&queryList); err != nil {
		log.Println("c.ShouldBindQuery queryList: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	items, err := pkg.Server.storageRepository.List(
		c.Request.Context(),
		queryList.Offset,
		queryList.Limit,
		queryList.Search,
	)
	if err != nil {
		log.Println("storageRepository.List: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
