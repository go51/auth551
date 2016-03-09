package auth551

import (
	"errors"
)

type Auth struct {
	config *Config
}

var authInstance *Auth

type Config struct {
	MasterKey     string       `json:"master_key"`
	CookieKeyName string       `json:"cookie_key_name"`
	Google        ConfigOAuth2 `json:"google"`
	Twitter       ConfigOAuth1 `json:"twitter"`
	Facebook      ConfigOAuth2 `json:"facebook"`
}

type ConfigOAuth2 struct {
	Vendor       string   `json:"vendor"`
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectUrl  string   `json:"redirect_url"`
	Scope        []string `json:"scope"`
	AuthUrl      string   `json:"auth_url"`
	TokenUrl     string   `json:"token_url"`
}

type ConfigOAuth1 struct {
	Vendor            string `json:"vendor"`
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	RedirectUrl       string `json:"redirect_url"`
	AuthorizeTokenUrl string `json:"authorize_token_url"`
	RequestTokenUrl   string `json:"request_token_url"`
	AccessTokenUrl    string `json:"access_token_url"`
}

type AuthVendor int

const (
	VENDOR_GOOGLE AuthVendor = iota
	VENDOR_TWITTER
	VENDOR_FACEBOOK
)

type AuthInstance interface {
	AuthCodeUrl() (url, token, secret string)
	TokenExchange(code, token, secret string) (*AccessToken, error)
	UserInfo(*AccessToken) (*AccountInformation, error)
}

type TweetInstance interface {
	Tweet(accessToken *AccessToken, tweet *TweetModel) error
}

func (av AuthVendor) String() string {
	switch av {
	case VENDOR_GOOGLE:
		return "google"
	case VENDOR_TWITTER:
		return "twitter"
	case VENDOR_FACEBOOK:
		return "facebook"
	default:
		return "unknown"
	}
}

func AuthVendorCode(vendor string) (AuthVendor, error) {
	switch vendor {
	case VENDOR_GOOGLE.String():
		return VENDOR_GOOGLE, nil

	case VENDOR_TWITTER.String():
		return VENDOR_TWITTER, nil

	case VENDOR_FACEBOOK.String():
		return VENDOR_FACEBOOK, nil

	default:
		return 0, errors.New("unknown vendor.")
	}
}

func Load(config *Config) *Auth {
	if authInstance != nil {
		return authInstance
	}

	authInstance = &Auth{
		config: config,
	}

	return authInstance
}

func (a *Auth) Auth(vendor AuthVendor) AuthInstance {
	switch vendor {
	case VENDOR_GOOGLE:
		return LoadGoogle(a.config.Google)
	case VENDOR_TWITTER:
		return LoadTwitter(a.config.Twitter)
	case VENDOR_FACEBOOK:
		return LoadFacebook(a.config.Facebook)
	default:
		return nil
	}
}

func (a *Auth) AuthCodeUrl(vendor AuthVendor) (url, token, secret string) {
	auth := a.Auth(vendor)
	if auth == nil {
		return "", "", ""
	}

	return auth.AuthCodeUrl()
}

func (a *Auth) UserInfoExchange(vendor AuthVendor, code, token, secret string) (*AccessToken, *AccountInformation, error) {
	auth := a.Auth(vendor)

	accessToken, err := auth.TokenExchange(code, token, secret)
	if err != nil {
		return nil, nil, err
	}

	accountInformation, err := auth.UserInfo(accessToken)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, accountInformation, nil

}

func (a *Auth) MasterKey() string {
	return a.config.MasterKey
}

func (a *Auth) CookieKeyName() string {
	return a.config.CookieKeyName
}
