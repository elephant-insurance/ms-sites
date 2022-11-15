package bc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SayOK returns a 200
// Add a HEAD route to this endpoint for heartbeat checks
func SayOK(c *gin.Context) {
	RenderJSONResponse(c, http.StatusOK, nil)
}

func SayBadRequest(c *gin.Context) {
	RenderJSONResponse(c, http.StatusBadRequest, nil)
}

func SayInternalServerError(c *gin.Context) {
	RenderJSONResponse(c, http.StatusInternalServerError, nil)
}

func SayNotFound(c *gin.Context) {
	RenderJSONResponse(c, http.StatusNotFound, nil)
}

func SayPreconditionFailed(c *gin.Context) {
	RenderJSONResponse(c, http.StatusPreconditionFailed, nil)
}
