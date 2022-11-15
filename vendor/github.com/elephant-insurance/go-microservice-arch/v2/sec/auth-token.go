package sec

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	grq "github.com/parnurzeal/gorequest"
)

// Check token end-point
const pathCheckToken = "check-token"

// Update user information
const pathPostUpdateUserInfo = "user-info"

// Request wait timeout
const idsRequestTimeoutInMS = 45000

type authTokenService interface {
	validateAndGetClaims(c msrqc.Context, authToken *AuthToken) (*AccessToken, bool, error)
	updateUserDetails(c msrqc.Context, userInfo *UserInfo, authToken *AuthToken) (bool, error)
}

func UpdateUserDetails(c msrqc.Context, userInfo *UserInfo, authToken *AuthToken) (bool, error) {
	return theAuthTokenService.updateUserDetails(c, userInfo, authToken)
}

func ValidateAndGetClaims(c msrqc.Context, authToken *AuthToken) (*AccessToken, bool, error) {
	return theAuthTokenService.validateAndGetClaims(c, authToken)
}

type realAuthTokenServiceType struct{}

// AuthToken => Auth token validation model
type AuthToken struct {
	AccessToken string `json:"-"`
	TokenType   string `json:"-"`

	Type   string `json:"type"`
	Error  string `json:"error"`
	Brand  string `json:"Brand"`
	Domain string `json:"Domain"`

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

// validateAndGetClaims => Validate the token ad extract scopes from the token
func (authTokenService *realAuthTokenServiceType) validateAndGetClaims(c msrqc.Context, obj *AuthToken) (*AccessToken, bool, error) {
	lw := log.ForFunc(c)
	lw.Debug("called")
	if obj == nil || obj.AccessToken == "" {
		e := errors.New("token object can't be blank or null")
		lw.WithError(e).Warn("invalid arguments")
		return nil, false, e
	}

	requestURL := fmt.Sprintf("%v/%v?brandkey=%v", msLoginBaseURL, pathCheckToken, obj.Brand)

	resp, body, errs := getRequest(c, requestURL).
		Set("authorization", fmt.Sprintf("%v %v", obj.TokenType, obj.AccessToken)).
		End()

	if len(errs) > 0 {
		lw.WithError(errs[0]).Error("error from ms-login")
		return nil, false, errs[0]
	}

	// 200: OK token is valid, 400: Token is invalid
	if body == "" || resp.StatusCode != 200 {
		lw.WithHTTPStatus(resp.StatusCode).Error("error from ms-login -" + body)
		return nil, false, errors.New("invalid token")
	}

	byteBody := []byte(body)

	var accessTok AccessToken
	if err := json.Unmarshal(byteBody, &accessTok); err != nil {
		lw.WithError(err).Error("error unmarshaling request body")
		return nil, false, nil
	}

	accessTok.LoadDataFromScopes()

	return &accessTok, !accessTok.IsExpired(), nil
}

// updateUserDetails => Update user details on UAA i.e. username, email id
func (authTokenService *realAuthTokenServiceType) updateUserDetails(c msrqc.Context, userInfo *UserInfo, obj *AuthToken) (bool, error) {
	lw := log.ForFunc(c)
	if obj == nil || obj.AccessToken == "" {
		e := errors.New("token object can't be blank or null")
		return false, e
	}

	requestURL := fmt.Sprintf("%v/%v", msLoginBaseURL, pathPostUpdateUserInfo)

	jsonReq, err := json.Marshal(userInfo)
	if err != nil {
		return false, err
	}

	resp, body, errs := postRequest(c, requestURL).
		Set("authorization", fmt.Sprintf("%v %v", obj.TokenType, obj.AccessToken)).
		Send(string(jsonReq)).
		End()

	if len(errs) > 0 {
		lw.WithError(errs[0]).Error("AuthToken.updateUserDetails error from ms-login")
		return false, errors.New("error from service")
	}

	// 201: OK, 400: Token is invalid
	if body == "" || (resp.StatusCode != 201 && resp.StatusCode != 400) {
		if resp.StatusCode != 201 && resp.StatusCode != 400 {
			lw.WithHTTPStatus(resp.StatusCode).Error("AuthToken.ValidateAndGetClaims bad response from ms-login")
		}

		return false, errors.New("invalid token")
	}

	return body != "", nil
}

// getRequest => send get type request to ms-login.
func getRequest(c context.Context, url string) *grq.SuperAgent {
	ctx := msrqc.New(c)
	idsReq := msrqc.GetRequestWithHeaders(ctx, url)

	idsReq.Timeout(idsRequestTimeoutInMS * time.Millisecond)
	idsReq.Set("Connection", "close")

	return idsReq
}

// postRequest => send post request to ms-login.
func postRequest(c context.Context, url string) *grq.SuperAgent {
	ctx := msrqc.New(c)
	idsReq := msrqc.PostRequestWithHeaders(ctx, url)

	idsReq.Timeout(idsRequestTimeoutInMS * time.Millisecond)
	idsReq.Set("Connection", "close")

	return idsReq
}
