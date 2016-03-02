package auth551

import (
	"encoding/json"
	xoauth2 "golang.org/x/oauth2"
)

type authFacebook struct {
	config *xoauth2.Config
}

var authFacebookInstance *authFacebook

func LoadFacebook(config ConfigOAuth2) *authFacebook {
	if authFacebookInstance != nil {
		return authFacebookInstance
	}

	authFacebookInstance = &authFacebook{
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

	return authFacebookInstance
}

func (a *authFacebook) AuthCodeUrl() (url, token, secret string) {
	return a.config.AuthCodeURL(""), "", ""

}

func (a *authFacebook) TokenExchange(code, token, secret string) (*AccessToken, error) {
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

func (a *authFacebook) UserInfo(accessToken *AccessToken) (*AccountInformation, error) {
	token := &xoauth2.Token{
		AccessToken:  accessToken.AccessToken,
		TokenType:    accessToken.TokenType,
		RefreshToken: accessToken.RefreshToken,
		Expiry:       accessToken.Expiry,
	}

	client := a.config.Client(nil, token)

	res, err := client.Get("https://graph.facebook.com/v2.5/me?fields=picture,id,name,email")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	account := myFacebookAccount{}
	err = json.NewDecoder(res.Body).Decode(&account)
	if err != nil {
		return nil, err
	}

	accountInformation := &AccountInformation{
		Id:      account.Id,
		Name:    account.Name,
		Email:   account.Email,
		Picture: account.Picture.Data.Url,
	}

	return accountInformation, nil
}

type myFacebookAccount struct {
	Id      string                   `json:"id"`
	Name    string                   `json:"name"`
	Email   string                   `json:"email"`
	Picture myFacebookAccountPicture `json:"picture"`
}

type myFacebookAccountPicture struct {
	Data myFacebookAccountPictureData `json:"data"`
}

type myFacebookAccountPictureData struct {
	IsSilhouette bool   `json:"is_silhouette"`
	Url          string `json:"url"`
}
