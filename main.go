package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/elephant-insurance/go-microservice-arch/v2/cors"
	"github.com/elephant-insurance/go-microservice-arch/v2/dig"
	"github.com/elephant-insurance/go-microservice-arch/v2/glog"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/go-microservice-arch/v2/sec"

	"github.com/elephant-insurance/ms-sites/app/cfg"
	"github.com/elephant-insurance/ms-sites/app/controllers"
	"github.com/elephant-insurance/ms-sites/app/routes"
	"github.com/elephant-insurance/ms-sites/app/services"
)

// main() should be the same for every microservice
func main() {
	// setup the app
	cfg.Initialize(nil)
	log.Initialize(cfg.Config.RequiredConfig, &cfg.Config.Logging)

	c := msrqc.New(context.Background())
	lw := log.ForFunc(c)

	// initialize model, service, and other packages
	setupApplicationPackages(c)

	// setup the web service
	lw.Debug("initializing gin...")
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	lw.Debug("gin initialized")

	// Add recovery handler FIRST so that it runs LAST
	g.Use(glog.RecoveryHandler())
	lw.Debug("recovery initialized")

	// uncomment if we're using compressed requests
	// g.Use(gzip.Gzip(cfg.Config.RequiredConfig))
	// lw.Debug("gzip initialized")

	dig.Initialize(cfg.Config.RequiredConfig, &cfg.Config.Diagnostics, g)
	lw.Debug("dig initialized")
	setupDiagnostics(c)

	sec.Initialize(cfg.Config.RequiredConfig, &cfg.Config.Security)

	g.Use(glog.New(cfg.Config.RequiredConfig))
	lw.Debug("glog initialized")

	g.Use(cors.New(cfg.Config.RequiredConfig))
	lw.Debug("cors initialized")

	router := routes.Initialize(cfg.Config.RequiredConfig, g)
	lw.Debug("routes initialized, listening...")

	router.Listen()
}

// setupApplicationPackages is the place to initialize any app-specific packages
// that require config settings or other work here in main.go.
// Use this to keep the body of main() the same for all microservices.
func setupApplicationPackages(c context.Context) {
	lw := log.ForFunc(c)
	bs := services.NewBlobService(cfg.Config.StorageAccountName, cfg.Config.StorageAccountKey, "singlesearch")
	services.Blob = bs
	lw.Debug(`application package initialization complete`)
}

// setupDiagnostics is the place to initialize any diagnostic tests or package stats.
// Use this to keep the body of main() the same for all microservices.
func setupDiagnostics(c context.Context) {
	lw := log.ForFunc(c)
	dig.AddPackageStats("default-log", log.Diagnostics)
	dig.AddPackageStats("cache-info", controllers.Diagnostics)
	dig.AddDiagnosticTest("storage-container-connection-test", services.TestContainer)
	// uncomment if we're using compressed requests
	// dig.AddPackageStats(`gzip`, gzip.Diagnostics)
	lw.Debug(`diagnostic tests and stats methods initialized`)
}
