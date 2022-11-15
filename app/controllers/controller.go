package controllers

import (
	"path/filepath"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"
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

	if downloadData, found := services.AllContents[docPath]; found {
		fileExtension := filepath.Ext(docPath)

		mimeType := services.DetectMimeType(strings.Split(fileExtension, ".")[1])
		// mimeType := http.DetectContentType(downloadData)

		c.Header("Content-Type", mimeType)
		// c.Header("Content-Disposition", "attachment")
		// c.Data(http.StatusOK, mimeType, downloadData)
		// resource := bytes.NewBuffer(downloadData)
		// _, _ = io.Copy(c.Writer, resource)
		c.Writer.Write(downloadData)
	}
	// resource := c.Param("resource")

	// for _, project := range services.Projects {
	// 	if project.Name == pName {
	// 		//serve from here
	// 	}
	// }

	// c.JSON(200, gin.H{})

	// retrieveTimer := dig.StartClientTiming(c, uf.Pointer.ToString(`retrieve-doc`), nil)
	// if doc, err := services.RetrieveDocument(c, docID); err != nil {
	// 	retrieveTimer.Stop(http.StatusNotFound)
	// 	bc.SayNotFound(c)
	// } else {
	// 	retrieveTimer.Stop(http.StatusOK)
	// 	bc.RenderJSONResponse(c, http.StatusOK, doc)
	// }
	lw.Debug(`complete`)
}
