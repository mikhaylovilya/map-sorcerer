package config

import (
	"log" //nolint:depguard
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvDebug string = "debug"
	EnvProd  string = "prod"
)

type Config struct {
	Env        string     `yaml:"env"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Port        uint64        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func Parse() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("No CONFIG_PATH found")
		// return nil, errors.New("No CONFIG_PATH found")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("No config file %s was found", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Failed to parse config: %s", err)
	}

	return &config
}
