package auth551

import "golang.org/x/oauth2"

type Auth struct {
	config *Config
}

var authInstance *Auth

type Config struct {
	Form   ConfigForm  `json:"form"`
	Google ConfigOAuth `json:"google"`
}

type ConfigForm struct {
	LoginId string `json:"loginId"`
}

type ConfigOAuth struct {
	Vendor       string   `json:"vendor"`
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectUrl  string   `json:"redirect_url"`
	Scope        []string `json:"scope"`
	AuthUrl      string   `json:"auth_url"`
	TokenUrl     string   `json:"token_url"`
}

type AuthVendor string

const (
	VENDOR_GOOGLE AuthVendor = "google"
)

func Load(config *Config) *Auth {
	if authInstance != nil {
		return authInstance
	}

	authInstance = &Auth{
		config: config,
	}

	return authInstance
}

func (a *Auth) authConfig(vendor AuthVendor) *oauth2.Config {
	var config ConfigOAuth

	switch vendor {
	case VENDOR_GOOGLE:
		config = a.config.Google
	default:
		return nil
	}

	authConfig := &oauth2.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectUrl,
		Scopes:       config.Scope,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthUrl,
			TokenURL: config.TokenUrl,
		},
	}

	return authConfig

}

func (a *Auth) AuthCodeUrl(vendor AuthVendor) string {
	authConfig := a.authConfig(vendor)

	return authConfig.AuthCodeURL("", oauth2.SetAuthURLParam("access_type", "offline"))
}

func (a *Auth) TokenExchange(vendor AuthVendor, code string) *oauth2.Token {
	authConfig := a.authConfig(vendor)

	token, err := authConfig.Exchange(nil, code)
	if err != nil {
		return nil
	}

	return token
}
