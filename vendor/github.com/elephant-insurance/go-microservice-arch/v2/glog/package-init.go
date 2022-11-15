package glog

import (
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/gin-gonic/gin"
)

// New sets up Gin-logging and returns a Gin handlerfunc for Gin to use
func New(requiredConfig cfg.Configurator) gin.HandlerFunc {
	appPrefix = requiredConfig.GetAppAbbreviation()
	return new(appPrefix)
}

// new wraps initialization for testing
func new(appAbbreviation string) gin.HandlerFunc {
	appPrefix = appAbbreviation
	return gin.HandlerFunc(serveHTTP)
}
