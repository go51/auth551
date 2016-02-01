package auth551

type Auth struct {}

var authInstance *Auth

func Load() *Auth {
	if authInstance != nil {
		return authInstance
	}

	authInstance = &Auth{}

	return authInstance
}
