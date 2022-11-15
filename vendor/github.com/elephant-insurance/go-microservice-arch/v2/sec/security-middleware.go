package sec

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
)

var (
	accessKey1, accessKey2, oktaBaseAddress, oktaClientID, msLoginBaseURL string
	initialized, bypassInDev, oktaMode, msLoginMode                       = false, false, false, false
	theAuthTokenService                                                   authTokenService
)

const (
	apiKeyHeaderKey          string = `apikey`
	aud                      string = `api://default`
	AuthHeaderKey            string = `Authorization`
	bypassWarning            string = `BYPASSING security in dev/test environment`
	oktaAuthEndPoint         string = `/oauth2/default`
	userClaims               string = `UserClaims`
	ContextKeyAccessToken    string = `AccessToken`
	PolicyScopePrefix        string = "guidewire.edge.policy."
	AccountScopePrefix       string = "guidewire.edge.account."
	LegacyAccountScopePrefix string = "guidewire.edge.legacyaccount."
	ScopePostfix             string = ".all"
)

// Initialize sets up any shared variables known at startup, which can be passed in from other packages (e.g., configuration).
// Initialize will not crash, but routes will crash if they are added with a security handler, if the package is not initialized.
func Initialize(requiredConfig cfg.Configurator, s *Settings) {
	lw := log.ForFunc(context.Background())

	if s != nil && s.OktaBaseAddress != `` && s.OktaClientID != `` {
		oktaBaseAddress = s.OktaBaseAddress
		oktaClientID = s.OktaClientID
		initialized = true
		oktaMode = true
	} else if s != nil && s.AccessKey1 != `` && s.AccessKey2 != `` {
		accessKey1 = s.AccessKey1
		accessKey2 = s.AccessKey2
		initialized = true
	} else if s != nil && s.MsLoginBaseURL != `` {
		msLoginBaseURL = s.MsLoginBaseURL
		initialized = true
		msLoginMode = true
	} else {
		lw.Warn(`no Okta settings or access keys found, security package is NOT initialized`)
	}

	theAuthTokenService = &realAuthTokenServiceType{}

	if s != nil && s.ByPassInDev != nil && *s.ByPassInDev {
		env := requiredConfig.GetEnvironment()
		if env.Equals(&enum.ServiceEnvironment.Development.ID) || env.Equals(&enum.ServiceEnvironment.Testing.ID) {
			lw.Warn(bypassWarning)
			bypassInDev = true
			theAuthTokenService = &authTokenServiceMock{}
		}
	}
}

// AuthorizeUserForHandler => AuthorizeUserForHandler authorize the incoming request and if there is no Key present is request then return response from here.
func AuthorizeUserForHandler(handler func(c *gin.Context)) func(c *gin.Context) {
	lw := log.ForFunc(context.Background())

	if !initialized {
		lw.Error("attempted to load secure route but sec package not initialized, returning 401")
		return func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
		}
	}

	if bypassInDev {
		lw.Debug(`returning BYPASS handler`)
		return func(c *gin.Context) {
			lwinner := log.ForFunc(c)
			lwinner.Warn(bypassWarning)
			handler(c)
		}
	}

	lw.Debug(`returning REAL handler`)

	if oktaMode {
		return func(c *gin.Context) {
			lw := log.ForFunc(c)
			var err error
			var result *jwtverifier.Jwt
			authHeaderValue := c.GetHeader(AuthHeaderKey)
			token := strings.Split(authHeaderValue, " ")
			lw.WithConsoleFields("token", token, AuthHeaderKey).Debug("checking token")

			if len(token) == 2 {
				if strings.EqualFold(token[0], "Bearer") && strings.Trim(token[1], " ") != "" {
					tv := map[string]string{}
					tv["cid"] = oktaClientID
					tv["aud"] = aud
					jv := jwtverifier.JwtVerifier{
						Issuer:           oktaBaseAddress + oktaAuthEndPoint,
						ClaimsToValidate: tv,
					}
					result, err = jv.New().VerifyAccessToken(token[1])

					if err == nil {
						var claims OktaClaims
						if byteClaims, jsonerr := json.Marshal(result.Claims); jsonerr == nil {
							if unmarshalErr := json.Unmarshal(byteClaims, &claims); unmarshalErr == nil {
								claims.Token = authHeaderValue
								c.Set(userClaims, claims)
								lw.Debug("access granted")
								handler(c)
								return
							}
						}
					}
				}
				lw.IfError(err).Info("not able to validate token")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
				return
			}

			lw.Debug("access not required")
			handler(c)
		}
	} else if msLoginMode {
		return AuthorizeUsingMsLogin(handler)
	}

	return func(c *gin.Context) {
		lw := log.ForFunc(c)
		key := c.Query("key")
		if key == "" {
			key = c.GetHeader(apiKeyHeaderKey)
		}

		lw.WithConsoleFields("key", key, AuthHeaderKey, accessKey1).Debug("checking key")

		if !(key == accessKey1 || key == accessKey2) {
			lw.Error("Invalid access attempt")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
			return
		}

		lw.Debug("access granted")
		handler(c)
	}
}

// AuthorizeUsingMsLogin => Business logic for Authorizing from ms-login via UAA.
func AuthorizeUsingMsLogin(handler func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		lw := log.ForFunc(c)
		lw.Debug("AuthorizeUsingMsLogin")
		auth := c.GetHeader(AuthHeaderKey)
		brand := msrqc.GetTransactionBrand(c)
		brandKey := ""
		if brand != nil {
			brandKey = brand.ToIDString()
		}
		if auth != "" {
			lw.WithConsoleFields("auth-header", auth).Debug("Authorizer.ServeHttp found authorization header")
			headerParts := strings.Split(auth, " ")
			if len(headerParts) > 1 {
				//  Try to get roles in action
				authToken := &AuthToken{AccessToken: headerParts[1], TokenType: headerParts[0], Brand: brandKey}
				accessTkn, isValid, _ := theAuthTokenService.validateAndGetClaims(c, authToken)
				if isValid || accessTkn != nil {
					contextSet(c, ContextKeyAccessToken, *accessTkn)
					handler(c)
					return
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Not Authorized"})
	}
}

func contextSet(c msrqc.Context, key string, value interface{}) {
	msrqc.NamespaceSet(c, &msrqc.NamespaceKeySec, &key, value, true)
}

func contextGet(c msrqc.Context, key string) (val interface{}, exists bool) {
	return msrqc.NamespaceGet(c, &msrqc.NamespaceKeySec, &key)
}
