package config

func init() {
	config.SetDefault("server.network", "tcp")
	config.SetDefault("server.address", ":9090")
	config.SetDefault("server.interceptors", []string{
		"ctxtags",
		"opentracing",
		"prometheus",
		"zap",
		"payload",
		"auth",
		"validator",
		"recovery",
	})
}

type server struct{}

// Server returns the collection of the server config.
var Server = &server{}

// Network returns the network of the gRPC server.
func (s *server) Network() string {
	return config.GetString("server.network")
}

// Address returns the address of the gRPC server.
func (s *server) Address() string {
	return config.GetString("server.address")
}

// Interceptors returns the interceptors of the gRPC server.
func (s *server) Interceptors() []string {
	return config.GetStringSlice("server.interceptors")
}
