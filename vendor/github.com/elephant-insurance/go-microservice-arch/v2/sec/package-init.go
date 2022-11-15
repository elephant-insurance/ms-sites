package sec

type Settings struct {
	AccessKey1      string `yaml:"AccessKey1" config:"optional"`
	AccessKey2      string `yaml:"AccessKey2" config:"optional"`
	ByPassInDev     *bool  `yaml:"BypassInDev" config:"optional"`
	OktaBaseAddress string `yaml:"OktaBaseAddress" config:"optional"`
	OktaClientID    string `yaml:"OktaClientID" config:"optional"`
	MsLoginBaseURL  string `yaml:"MsLoginBaseURL" config:"optional"`
}

type OktaClaims struct {
	Audience  string   `json:"aud"`
	ClientID  string   `json:"cid"`
	ExpiredAt int      `json:"exp"`
	IssuedAt  int      `json:"iat"`
	Issuer    string   `json:"iss"`
	Jti       string   `json:"jti"`
	Scope     []string `json:"scp"`
	Subject   string   `json:"sub"`
	Version   int      `json:"ver"`
	Token     string   `json:"token"`
}

type UserInfo struct {
	Username      string
	Email         string
	FirstName     string
	LastName      string
	BrandKey      string
	AccountDomain string
}

// GetEnvironmentSetting implements the Overridable interface for the configuration package
// Given a field name, it returns the name of the corresponding environment variable
func (s Settings) GetEnvironmentSetting(fieldName string) string {
	return `MSVCSEC_` + fieldName
}
