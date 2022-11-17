package routes

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/routes"
	c "github.com/elephant-insurance/ms-sites/app/controllers"
)

var appRouter *routes.Router

// Initialize initializes routes for the app
func Initialize(requiredConfig cfg.Configurator, g *gin.Engine) *routes.Router {
	log.ForFunc(context.Background()).Debug("loading routes")

	appRouter = routes.New(requiredConfig, g)

	appRouter.GET(routeNameGetDocument, pathGetDocument, c.HandleGetDocument)
	appRouter.GET(routeNameGetDocument, pathGetIndex, c.HandleGetDocument)

	return appRouter
}
