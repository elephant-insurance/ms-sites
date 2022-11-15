package routes

import (
	"testing"

	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/routes"

	enum "github.com/elephant-insurance/enumerations/v2"

	"github.com/gin-gonic/gin"
)

var routeTests = []routes.RouteTest{}

func TestRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	g := gin.New()
	rc := cfg.RequiredConfig{
		Environment: enum.ServiceEnvironment.Testing.ID,
	}
	testRC := cfg.NewTestConfigurator(rc)
	r := Initialize(testRC, g)
	r.FinalizeForTest(testRC)

	for _, thisTest := range routeTests {
		thisTest.Run(r, g, t)
	}
}
