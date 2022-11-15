package sec

import (
	"strings"
	"time"

	"github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
)

// AccessToken => Access token details
type AccessToken struct {
	Brand               string `json:"-"`
	AccountDomain       string `json:"-"`
	PolicyNumber        string `json:"-"`
	AccountNumber       string `json:"-"`
	LegacyAccountNumber string `json:"-"`

	Scope    []string `json:"scope"`
	Audience []string `json:"aud"`
	UserID   string   `json:"user_id"`
	Subject  string   `json:"sub"`
	UserName string   `json:"user_name"`

	Email string `json:"email"`
	Jti   string `json:"jti"`
	Zid   string `json:"zid"`

	Origin    string `json:"origin"`
	Revocable bool   `json:"revocable"`
	AuthTime  *int64 `json:"auth_time"`
	ExpTime   *int64 `json:"exp"`
	Iat       *int64 `json:"iat"`
}

// LoadDataFromScopes => Loads data into properties from Scope section
// i.e. PolicyNumber
func (accessToken *AccessToken) LoadDataFromScopes() {
	accessToken.Brand = "elephant"

	if accessToken.Zid != "uaa" {
		accessToken.Brand = accessToken.Zid
	}

	if len(accessToken.Scope) > 0 {
		for _, scope := range accessToken.Scope {
			if strings.Contains(scope, PolicyScopePrefix) {
				accessToken.PolicyNumber = strings.Replace(strings.Replace(scope, PolicyScopePrefix, "", 1), ScopePostfix, "", 1)
			} else if strings.Contains(scope, AccountScopePrefix) {
				accessToken.AccountNumber = strings.Replace(strings.Replace(scope, AccountScopePrefix, "", 1), ScopePostfix, "", 1)
			} else if strings.Contains(scope, LegacyAccountScopePrefix) {
				accessToken.LegacyAccountNumber = strings.Replace(strings.Replace(scope, LegacyAccountScopePrefix, "", 1), ScopePostfix, "", 1)
			}
		}

		if accessToken.LegacyAccountNumber == "" {
			accessToken.LegacyAccountNumber = accessToken.AccountNumber
		}
	}
}

// GetExpirationTime : gets the expiration time from the token
func (accessToken *AccessToken) GetExpirationTime() *time.Time {
	var result time.Time

	if accessToken.ExpTime != nil {
		result = time.Unix(*accessToken.ExpTime, 0)
	}

	return &result
}

// IsExpired : checks if the access token is expired
func (accessToken *AccessToken) IsExpired() bool {
	if accessToken.ExpTime == nil {
		return true
	}

	duration := time.Since(*accessToken.GetExpirationTime())

	return int64(duration.Seconds()) >= 2
}

// TokenFromContext : extracts token from the context
func TokenFromContext(c msrqc.Context) *AccessToken {
	lw := log.ForFunc(c)
	if t, exists := contextGet(c, ContextKeyAccessToken); exists && t != nil {
		if accessTkn, ok := t.(AccessToken); ok {
			return &accessTkn
		}
	}

	lw.Warn("TokenFromContext returning NIL AccessToken")
	return nil
}

// IsRoleExist : checks if the access token has roles.
func (accessToken *AccessToken) IsRoleExist(roles []enumerations.AgentRoleID) bool {
	for _, role := range roles {
		for _, scope := range accessToken.Scope {
			if roleEnumItem := enumerations.AgentRole.ByIDString(scope); roleEnumItem != nil && role.Equals(&roleEnumItem.ID) {
				return true
			}
		}
	}
	return false
}
