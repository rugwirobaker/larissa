package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	validator "gopkg.in/go-playground/validator.v9"
)

const defaultConfigFile = "larissa.toml"

// Config provides configuration values for all components
type Config struct {
	Port        string `envconfig:"LARISSA_PORT"` // http port
	StorageType string `validate:"required" envconfig:"LARISSA_STORAGE_TYPE"`
	Storage     *StorageConfig
}

func defaultConfig() *Config {
	disk := &DiskConfig{"data"}
	store := &StorageConfig{disk}
	return &Config{
		Port:        "3000",
		StorageType: "disk",
		Storage:     store,
	}
}

// Load loads the config from a file.
// If file is not present returns default config
func Load(configFile string) (*Config, error) {
	// User explicitly specified a config file
	if configFile != "" {
		return ParseConfigFile(configFile)
	}

	// There is a config in the current directory
	if fi, err := os.Stat(defaultConfigFile); err == nil {
		return ParseConfigFile(fi.Name())
	}

	// Use default values
	log.Println("Running dev mode with default settings, consult config when you're ready to run in production")
	return defaultConfig(), nil
}

// ParseConfigFile parses the given file into an larissa config struct
func ParseConfigFile(configFile string) (*Config, error) {
	var config Config
	// attempt to read the given config file
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}
	// validate all required fields have been populated
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	return &config, nil
}

func validateConfig(config Config) error {
	validate := validator.New()
	err := validate.StructExcept(config, "Storage")
	if err != nil {
		return err
	}
	switch config.StorageType {
	case "disk":
		return validate.Struct(config.Storage.Disk)
	default:
		return fmt.Errorf("storage type %s is unknown", config.StorageType)
	}
}
