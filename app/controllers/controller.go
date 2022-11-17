package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/dig"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
	"github.com/elephant-insurance/ms-sites/app/cfg"
	"github.com/elephant-insurance/ms-sites/app/services"
	"github.com/gin-gonic/gin"
)

func HandleGetDocument(c *gin.Context) {

	lw := log.ForFunc(c).Debug(`called`)
	var docPath string
	project := c.Param("project")
	doc := c.Param("document")

	if doc == "" {
		docPath = project + "/index.html"
	} else {
		docPath = project + `/` + doc
	}
	fileExtension := filepath.Ext(docPath)
	mimeType := services.DetectMimeType(strings.Split(fileExtension, ".")[1])

	retrieveTimer := dig.StartClientTiming(c, uf.Pointer.ToString(`retrieve-doc`), nil)
	// check if the doc is in cache
	if downloadData, cacheHit := services.FindInCache(docPath); cacheHit {
		lw.Debug("Cache hit")

		retrieveTimer.Stop(http.StatusOK)
		c.Header("Content-Type", mimeType)
		c.Writer.Write(downloadData)

	} else {
		lw.Debug("Cache miss")
		//download tarball and unzip and cache
		err, statusCode := cfg.BlobService.DownloadFiles(c, project)
		if err != nil {
			retrieveTimer.Stop(statusCode)
			if statusCode == http.StatusNotFound {
				retrieveTimer.Stop(http.StatusNotFound)
				c.Status(http.StatusNotFound)
			} else {
				c.Writer.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			{
				downloadData, _ := services.FindInCache(docPath)
				c.Header("Content-Type", mimeType)
				c.Writer.Write(downloadData)
				retrieveTimer.Stop(http.StatusOK)
			}
		}
	}

	lw.Debug(`complete`)
}
