package envconfig

const (
	defaultConfigFile       = "config-local.conf"
	defaultLocalEnvironment = "local"
	defaultLocalEnvFile     = ".env"
)

type Config struct {
	configFile  string
	environment string
}

type Option func(*Config)

func ConfigFile(configFile string) Option {
	return func(c *Config) {
		c.configFile = configFile
	}
}

func Environment(environment string) Option {
	return func(c *Config) {
		c.environment = environment
	}
}
