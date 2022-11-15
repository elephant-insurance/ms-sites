package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/elephant-insurance/go-microservice-arch/v2/cors"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/ms-sites/app/cfg"
	"github.com/elephant-insurance/ms-sites/app/routes"
	"github.com/gin-gonic/gin"
)

func TestInitialize(t *testing.T) {
	cfg.Initialize(nil)
	log.Initialize(cfg.Config.RequiredConfig, &cfg.Config.Logging)
	spew.Dump(cfg.Config)
	g := gin.New()
	cors.New(cfg.Config.RequiredConfig)
	router := routes.Initialize(cfg.Config.RequiredConfig, g)
	router.Listen()
}
