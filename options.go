package envconfig

const (
	defaultConfigFile = "config-local.conf"
)

type Config struct {
	configFile string
}

type Option func(*Config)

func ConfigFile(configFile string) Option {
	return func(c *Config) {
		c.configFile = configFile
	}
}
