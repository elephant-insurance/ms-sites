package sec

import "github.com/elephant-insurance/go-microservice-arch/v2/msrqc"

type authTokenServiceMock struct{}

// ValidateAndGetClaims => Validate the token ad extract scopes from the token
func (authTokenService *authTokenServiceMock) validateAndGetClaims(c msrqc.Context, obj *AuthToken) (*AccessToken, bool, error) {
	accessTok := &AccessToken{
		Brand:    "elephant",
		UserName: "shams13@gmail.com",
	}
	return accessTok, true, nil
}

// updateUserDetails => Update user details on UAA i.e. username, email id
func (authTokenService *authTokenServiceMock) updateUserDetails(c msrqc.Context, userInfo *UserInfo, obj *AuthToken) (bool, error) {
	return true, nil
}
