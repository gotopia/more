package config

func init() {
	config.SetDefault("gateway.address", ":8080")
}

type gateway struct{}

// Gateway returns the collection of the gateway config.
var Gateway = &gateway{}

// Address returns the address of the gRPC gateway.
func (g *gateway) Address() string {
	return config.GetString("gateway.address")
}
