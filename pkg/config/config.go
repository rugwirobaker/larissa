package config

const defaultConfigFile = "larissa.toml"

// Config provides configuration values for all components
type Config struct {
	Port     string `envconfig:"LARISSA_PORT"`              // http port
	RootPath string `envconfig:"LARISSA_DISK_STORAGE_ROOT"` // root path fo storage
}

func defaultConfig() *Config {
	return &Config{
		Port:     "3000",
		RootPath: "data",
	}
}

// Load loads the config from a file.
// If file is not present returns default config
func Load(configFile string) (*Config, error) {
	return defaultConfig(), nil
}

// ParseConfigFile parses the given file into an larissa config struct
func ParseConfigFile(configFile string) (*Config, error) {
	return nil, nil
}
