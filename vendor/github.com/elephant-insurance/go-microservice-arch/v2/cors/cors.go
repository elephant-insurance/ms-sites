package cors

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/gin-gonic/gin"
)

func New(requiredConfig cfg.Configurator) gin.HandlerFunc {
	lw := log.ForFunc(context.Background())

	headers := requiredConfig.GetAllowedHeaders()
	methods := requiredConfig.GetAllowedMethods()
	origins := requiredConfig.GetAllowedOrigins()
	exposed := requiredConfig.GetExposedHeaders()
	if len(origins) == 0 || origins[0] == `` {
		lw.Fatal(`invalid value for allowed origins`)
	}

	// we always allow these three:
	allowHeaders := standardAcceptHeaders

	for _, v := range headers {
		switch v {
		case AcceptHeaderXRequestedWith,
			AcceptHeaderOrigin,
			AcceptHeaderContentType,
			AcceptHeaderAccept,
			AcceptHeaderXAuthToken,
			AcceptHeaderXAPIVersionToken,
			AcceptHeaderAuthorization,
			enum.TXHeader.Brand.HeaderKey,
			enum.TXHeader.Domain.HeaderKey,
			enum.TXHeader.ID.HeaderKey,
			enum.TXHeader.Integrator.HeaderKey,
			enum.TXHeader.Source.HeaderKey,
			enum.TXHeader.Type.HeaderKey,
			enum.TXHeader.IPAddress.HeaderKey,
			enum.TXHeader.Instance.HeaderKey:
			// these are all standard
			continue
		default:
			allowHeaders = append(allowHeaders, v)
		}
	}

	exposedHeaders := standardExposeHeaders
	for _, v := range exposed {
		switch v {
		case ExposeHeaderContentLength,
			ExposeHeaderDate,
			enum.TXHeader.Brand.HeaderKey,
			enum.TXHeader.Domain.HeaderKey,
			enum.TXHeader.ID.HeaderKey,
			enum.TXHeader.Integrator.HeaderKey,
			enum.TXHeader.Source.HeaderKey,
			enum.TXHeader.Type.HeaderKey,
			enum.TXHeader.IPAddress.HeaderKey,
			enum.TXHeader.Instance.HeaderKey:
			// this is a standard exposed header, no need to add it again
			continue
		default:
			exposedHeaders = append(exposedHeaders, v)
		}
	}

	// we always allow these three:
	allowMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
	}

	for _, v := range methods {
		switch v {
		case http.MethodGet, http.MethodHead, http.MethodOptions:
			continue
		case http.MethodConnect, http.MethodDelete, http.MethodPatch, http.MethodPost, http.MethodPut, http.MethodTrace:
			allowMethods = append(allowMethods, v)
		default:
			lw.WithHTTPMethod(v).Fatal(`allowed method not recognized`)
		}
	}

	cfgstr := cors.Config{
		AllowMethods:     allowMethods,
		AllowHeaders:     allowHeaders,
		ExposeHeaders:    exposedHeaders,
		AllowWildcard:    true,
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(cfgstr)
}

var standardAcceptHeaders = []string{
	AcceptHeaderXRequestedWith,
	AcceptHeaderOrigin,
	AcceptHeaderContentType,
	AcceptHeaderAccept,
	AcceptHeaderXAuthToken,
	AcceptHeaderXAPIVersionToken,
	AcceptHeaderAuthorization,
	enum.TXHeader.Brand.HeaderKey,
	enum.TXHeader.Domain.HeaderKey,
	enum.TXHeader.ID.HeaderKey,
	enum.TXHeader.Integrator.HeaderKey,
	enum.TXHeader.Source.HeaderKey,
	enum.TXHeader.Type.HeaderKey,
	enum.TXHeader.IPAddress.HeaderKey,
	enum.TXHeader.Instance.HeaderKey,
}

var standardExposeHeaders = []string{
	ExposeHeaderContentLength,
	ExposeHeaderDate,
	enum.TXHeader.Brand.HeaderKey,
	enum.TXHeader.Domain.HeaderKey,
	enum.TXHeader.ID.HeaderKey,
	enum.TXHeader.Integrator.HeaderKey,
	enum.TXHeader.Source.HeaderKey,
	enum.TXHeader.Type.HeaderKey,
	enum.TXHeader.IPAddress.HeaderKey,
	enum.TXHeader.Instance.HeaderKey,
}
