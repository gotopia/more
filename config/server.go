package config

func init() {
	config.SetDefault("server.network", "tcp")
	config.SetDefault("server.address", ":9090")
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
