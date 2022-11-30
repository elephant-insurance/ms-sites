package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
	"github.com/elephant-insurance/go-microservice-arch/v2/dig"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
	"github.com/elephant-insurance/ms-sites/app/services"
	"github.com/gin-gonic/gin"
)

var (
	CacheHits *clicker.Clicker = &clicker.Clicker{}
	CacheMiss *clicker.Clicker = &clicker.Clicker{}

	timinigLabelCacheHits = uf.Pointer.ToString(`cache-hits`)
	timinigLabelCacheMiss = uf.Pointer.ToString(`cache-miss`)
)

func HandleGetDocument(c *gin.Context) {

	lw := log.ForFunc(c).Debug(`called`)
	var docPath string
	project := c.Param("project")
	doc := c.Param("document")
	if project == "" {
		c.Status(http.StatusNotFound)
		return
	}
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
		CacheHits.Click(1)
		retrieveTimer.Stop(http.StatusOK)
		c.Header("Content-Type", mimeType)
		c.Writer.Write(downloadData)

	} else {
		lw.Debug("Cache miss")
		CacheMiss.Click(1)
		//download tarball and unzip and cache
		err, statusCode := services.Blob.DownloadFiles(c, project)
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

func Diagnostics() map[string]interface{} {
	return map[string]interface{}{
		`cache-hits`: CacheHits.Clicks,
		`cache-miss`: CacheMiss.Clicks,
	}
}
