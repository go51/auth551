package auth551

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

func Load(config *Config) *Auth {
	if authInstance != nil {
		return authInstance
	}

	authInstance = &Auth{
		config: config,
	}

	return authInstance
}
