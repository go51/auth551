package auth551

import (
	"encoding/json"
	"github.com/mrjones/oauth"
)

type authTwitter struct {
	config   ConfigOAuth1
	consumer *oauth.Consumer
}

var authTwitterInstance *authTwitter

func LoadTwitter(config ConfigOAuth1) *authTwitter {
	if authTwitterInstance != nil {
		return authTwitterInstance
	}

	authTwitterInstance = &authTwitter{
		config: config,
		consumer: oauth.NewConsumer(
			config.ConsumerKey,
			config.ConsumerSecret,
			oauth.ServiceProvider{
				AuthorizeTokenUrl: config.AuthorizeTokenUrl,
				RequestTokenUrl:   config.RequestTokenUrl,
				AccessTokenUrl:    config.AccessTokenUrl,
			},
		),
	}

	return authTwitterInstance
}

func (a *authTwitter) AuthCodeUrl() (url, token, secret string) {
	requestToken, url, err := a.consumer.GetRequestTokenAndUrl(a.config.RedirectUrl)
	if err != nil {
		panic(err)
	}

	return url, requestToken.Token, requestToken.Secret
}

func (a *authTwitter) TokenExchange(code, token, secret string) (*AccessToken, error) {
	t := &oauth.RequestToken{
		Token:  token,
		Secret: secret,
	}

	authorizeToken, err := a.consumer.AuthorizeToken(t, code)
	if err != nil {
		return nil, err
	}

	accessToken := &AccessToken{
		AccessToken: authorizeToken.Token,
		TokenSecret: authorizeToken.Secret,
	}

	return accessToken, nil

}

func (a *authTwitter) UserInfo(accessToken *AccessToken) (*AccountInformation, error) {
	token := &oauth.AccessToken{
		Token:  accessToken.AccessToken,
		Secret: accessToken.TokenSecret,
	}

	res, err := a.consumer.Get(
		"https://api.twitter.com/1.1/account/verify_credentials.json",
		map[string]string{},
		token,
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	account := myTwitterAccount{}
	err = json.NewDecoder(res.Body).Decode(&account)
	if err != nil {
		return nil, err
	}

	accountInformation := &AccountInformation{
		Id:      account.Id,
		Name:    account.Name,
		Email:   "",
		Picture: account.Picture,
	}

	return accountInformation, nil
}

type myTwitterAccount struct {
	Id      string `json:"id_str"`
	Name    string `json:"name"`
	Picture string `json:"profile_image_url_https"`
}
