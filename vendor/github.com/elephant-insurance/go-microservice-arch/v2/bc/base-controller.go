package bc

import (
	"net/http"
	"strings"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/go-microservice-arch/v2/resp"
	"github.com/gin-gonic/gin"
)

// jsonContentType => response content type header value
const (
	jsonContentType = "application/json"
	htmlContentType = "text/html"
)

// RenderJSONResponse => Return response as JSON data
func RenderJSONResponse(ginCtx *gin.Context, status int, v interface{}) {
	ginCtx.Writer.Header().Set("Content-Type", jsonContentType)

	for key, val := range msrqc.HeadersFromContext(ginCtx) {
		if len(val) > 0 && val[0] != "" {
			ginCtx.Writer.Header().Add(key, val[0])
		}
	}

	RegisterResponse(status)

	ginCtx.JSON(status, v)
}

// RenderHTMLResponse => Return response as HTML with data
func RenderHTMLResponse(ginCtx *gin.Context, status int, v interface{}, htmlTemplateName string) {
	ginCtx.Writer.Header().Set("Content-Type", htmlContentType)

	for key, val := range msrqc.HeadersFromContext(ginCtx) {
		if len(val) > 0 && val[0] != "" {
			ginCtx.Writer.Header().Add(key, val[0])
		}
	}

	RegisterResponse(status)

	ginCtx.HTML(status, htmlTemplateName, v)
}

// HandleJSONResponse => handles http responses and returns response in JSON format
func HandleJSONResponse(ginCtx *gin.Context, resource *resp.ResponseMessageDto) {
	resource.LoadDefaults()
	ginCtx.Header("Date", time.Now().Format("01/02/2006 03:04:05 PM"))
	if isUnauthorized(resource) {
		RenderJSONResponse(ginCtx, http.StatusUnauthorized, resource)
	} else if resource.HasErrors() {
		RenderJSONResponse(ginCtx, http.StatusBadRequest, resource)
	} else if resource.HasValidations() {
		RenderJSONResponse(ginCtx, http.StatusUnprocessableEntity, resource)
	} else {
		RenderJSONResponse(ginCtx, http.StatusOK, resource)
	}
}

// isUnauthorized check if any of the error code is Unauthorized
func isUnauthorized(resource *resp.ResponseMessageDto) bool {
	if resource != nil && len(resource.ErrorMessages) > 0 {
		for _, msg := range resource.ErrorMessages {
			if msg == nil {
				continue
			}

			if strings.EqualFold(enum.CommonErrorCode.Unauthorized.Name, msg.Code) {
				return true
			}
		}
	}

	return false
}
