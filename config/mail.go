package config

func init() {
	config.SetDefault("mail.smtp.host", "localhost")
	config.SetDefault("mail.smtp.port", 587)
	config.SetDefault("mail.smtp.auth.type", "none")
	config.SetDefault("mail.smtp.auth.username", "")
	config.SetDefault("mail.smtp.auth.password", "")
}

type smtpAuth struct {
}

type smtp struct {
	Auth smtpAuth
}

type mail struct {
	SMTP smtp
}

// Mail returns the collection of the mail config.
var Mail = &mail{}

// Host returns the host of smtp server.
func (s *smtp) Host() string {
	return config.GetString("mail.smtp.host")
}

// Port returns the port of smtp server.
func (s *smtp) Port() int {
	return config.GetInt("mail.smtp.port")
}

// Type returns the type of authentication.
func (a *smtpAuth) Type() string {
	return config.GetString("mail.smtp.auth.type")
}

// Username returns the username of plain auth.
func (a *smtpAuth) Username() string {
	return config.GetString("mail.smtp.auth.username")
}

// Password returns the password of plain auth.
func (a *smtpAuth) Password() string {
	return config.GetString("mail.smtp.auth.password")
}
