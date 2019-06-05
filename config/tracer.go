package config

func init() {
	config.SetDefault("tracer.enabled", false)
	config.SetDefault("tracer.name", "jaeger")
}

type tracer struct {
	Agent agent
}

type agent struct {
}

// Tracer returns the collection of the tracer config.
var Tracer = &tracer{}

// Enabled checks whether the tracer is enabled.
func (t *tracer) Enabled() bool {
	return config.GetBool("tracer.enabled")
}

// Name returns the name of the tracer.
func (t *tracer) Name() string {
	return config.GetString("tracer.name")
}

// ServiceName returns the service_name of the tracer.
func (t *tracer) ServiceName() string {
	return config.GetString("tracer.service_name")
}

// Host returns the host of agent.
func (a *agent) Host() string {
	return config.GetString("tracer.agent.host")
}

// Port returns the port of agent.
func (a *agent) Port() int {
	return config.GetInt("tracer.agent.port")
}
