package auth551

import (
	xoauth2 "golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
)

type authGoogle struct {
	config *xoauth2.Config
}

var authGoogleInstance *authGoogle

func LoadGoogle(config ConfigOAuth2) *authGoogle {
	if authGoogleInstance != nil {
		return authGoogleInstance
	}

	authGoogleInstance = &authGoogle{
		config: &xoauth2.Config{
			ClientID:     config.ClientId,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.RedirectUrl,
			Scopes:       config.Scope,
			Endpoint: xoauth2.Endpoint{
				AuthURL:  config.AuthUrl,
				TokenURL: config.TokenUrl,
			},
		},
	}

	return authGoogleInstance
}

func (a *authGoogle) AuthCodeUrl() (url, token, secret string) {
	return a.config.AuthCodeURL("", xoauth2.SetAuthURLParam("access_type", "offline")), "", ""

}

func (a *authGoogle) TokenExchange(code, token, secret string) (*AccessToken, error) {
	exchangeAccessToken, err := a.config.Exchange(nil, code)
	if err != nil {
		return nil, err
	}

	accessToken := &AccessToken{
		AccessToken:  exchangeAccessToken.AccessToken,
		TokenType:    exchangeAccessToken.TokenType,
		RefreshToken: exchangeAccessToken.RefreshToken,
		Expiry:       exchangeAccessToken.Expiry,
	}

	return accessToken, nil
}

func (a *authGoogle) UserInfo(accessToken *AccessToken) (*AccountInformation, error) {
	token := &xoauth2.Token{
		AccessToken:  accessToken.AccessToken,
		TokenType:    accessToken.TokenType,
		RefreshToken: accessToken.RefreshToken,
		Expiry:       accessToken.Expiry,
	}

	client := a.config.Client(nil, token)

	auth2, err := oauth2.New(client)
	if err != nil {
		return nil, err
	}
	userInformation, err := auth2.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}

	accountInformation := &AccountInformation{
		Id:      userInformation.Id,
		Name:    userInformation.Name,
		Email:   userInformation.Email,
		Picture: userInformation.Picture,
	}

	return accountInformation, nil

}
