package dig

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
	"github.com/gin-gonic/gin"
)

// These functions wrap the Gin routing directives GET, POST, PUT, etc.
// They're needed because Gin's crappy router won't tell us which route it selected
// Each function captures the route name in a func closure at startup

// DELETE wraps a gin.DELETE call in a diagnostic timer
func DELETE(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodDelete, label, relativePath, handlers...)
}

// GET wraps a gin.GET call in a diagnostic timer
func GET(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodGet, label, relativePath, handlers...)
}

// HEAD wraps a gin.HEAD call in a diagnostic timer
func HEAD(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodHead, label, relativePath, handlers...)
}

// OPTIONS wraps a gin.OPTIONS call in a diagnostic timer
func OPTIONS(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodOptions, label, relativePath, handlers...)
}

// PATCH wraps a gin.PATCH call in a diagnostic timer
func PATCH(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodPatch, label, relativePath, handlers...)
}

// POST wraps a gin.POST call in a diagnostic timer
func POST(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodPost, label, relativePath, handlers...)
}

// PUT wraps a gin.PUT call in a diagnostic timer
func PUT(engine *gin.Engine, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return createHandler(engine, http.MethodPut, label, relativePath, handlers...)
}

func createHandler(engine *gin.Engine, method, label, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	// the controller is the last item in the handler chain
	// replace it with a wrapped version of the same func
	// this way we are only timing the controller action, and not
	// authentication, diagnostics, logging, routing, etc.
	hc := len(handlers) // handler count
	sep := "/"
	bp := engine.BasePath()
	if strings.HasSuffix(bp, "/") {
		sep = ""
	}
	fullPath := fmt.Sprintf("%v %v%v%v", method, bp, sep, relativePath)
	newHandlers := make([]gin.HandlerFunc, hc, len(handlers))
	copy(handlers[0:hc-1], newHandlers)
	ctrlr := handlers[hc-1]
	newctrlr := func(c *gin.Context) {
		timer := StartServiceTiming(c, &label, &fullPath)
		ctrlr(c)
		timer.Stop(c.Writer.Status())
	}
	newHandlers[hc-1] = newctrlr

	return engine.Handle(method, relativePath, newHandlers...)
}

// LastRouteTiming returns the timing info for the last service response timed by this package
// It is mainly useful for testing: Gin won't tell us which route it chose, but this will!
func LastRouteTiming() *timing.Timing {
	rtn := timing.ApplicationTimings.SelectCountFromLatest(selectRouteTiming, 1)
	if len(rtn) > 0 {
		return rtn[0]
	}

	return nil
}

func selectRouteTiming(obj interface{}) bool {
	t, ok := obj.(*timing.Timing)
	if !ok {
		return false
	}

	return t.Type == timing.TimingTypeWebService
}
