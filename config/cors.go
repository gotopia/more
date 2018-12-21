package config

func init() {
	config.SetDefault("cors.origins", []string{"localhost:4200"})
	config.SetDefault("cors.methods", []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
	})
	config.SetDefault("cors.headers", []string{
		"Authorization",
	})
}

type cors struct{}

// Cors returns the collection of the cors config.
var Cors = &cors{}

// Origins returns the allowed origins.
func (c *cors) Origins() []string {
	return config.GetStringSlice("cors.origins")
}

// Methods returns the allowed methods.
func (c *cors) Methods() []string {
	return config.GetStringSlice("cors.methods")
}

// Headers returns the allowed headers.
func (c *cors) Headers() []string {
	return config.GetStringSlice("cors.headers")
}
