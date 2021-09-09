package api

// ConfigWithDefaults sets default values and returns a new Config.
func ConfigWithDefaults() *Config {
	return &Config{}
}

// Config holds information that can be modified by config or command line flags.
type Config struct {
	ClusterPolicyUpdateKey string
}
