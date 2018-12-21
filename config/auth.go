package config

func init() {
	config.SetDefault("auth.type", "none")
}

type auth struct {
	Sign *sign
}
type sign struct{}

// Auth returns the collection of the auth config.
var Auth = &auth{
	Sign: &sign{},
}

// Type returns the type of auth.
func (a *auth) Type() string {
	return config.GetString("auth.type")
}

// Issuer returns the issuer of auth.
func (a *auth) Issuer() string {
	return config.GetString("auth.issuer")
}

// KeyDir returns the key directory of sign.
func (s *sign) KeyDir() string {
	return config.GetString("auth.sign.key_dir")
}

// Issuer returns the issuer of sign.
func (s *sign) Issuer() string {
	return config.GetString("auth.sign.issuer")
}
